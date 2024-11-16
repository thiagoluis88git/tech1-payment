package usecases

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
)

type GenerateQRCodePaymentUseCase interface {
	Execute(
		ctx context.Context,
		token string,
		qrOrder dto.QRCodeOrder,
		date int64,
		wg *sync.WaitGroup,
		ch chan bool,
	) (dto.QRCodeDataResponse, error)
}

type GenerateQRCodePaymentUseCaseImpl struct {
	repository        repository.QRCodePaymentRepository
	orderRepository   repository.OrderRepository
	paymentRepository repository.PaymentRepository
}

type FinishOrderForQRCodeUseCase interface {
	Execute(ctx context.Context, token string, form dto.ExternalPaymentEvent) error
}

type FinishOrderForQRCodeUseCaseImpl struct {
	repository        repository.QRCodePaymentRepository
	orderRepository   repository.OrderRepository
	paymentRepository repository.PaymentRepository
}

func NewGenerateQRCodePaymentUseCase(
	repository repository.QRCodePaymentRepository,
	orderRepository repository.OrderRepository,
	paymentRepository repository.PaymentRepository,
) GenerateQRCodePaymentUseCase {
	return &GenerateQRCodePaymentUseCaseImpl{
		repository:        repository,
		orderRepository:   orderRepository,
		paymentRepository: paymentRepository,
	}
}

func NewFinishOrderForQRCodeUseCase(
	repository repository.QRCodePaymentRepository,
	orderRepository repository.OrderRepository,
	paymentRepository repository.PaymentRepository,
) FinishOrderForQRCodeUseCase {
	return &FinishOrderForQRCodeUseCaseImpl{
		repository:        repository,
		orderRepository:   orderRepository,
		paymentRepository: paymentRepository,
	}
}

func (service *GenerateQRCodePaymentUseCaseImpl) Execute(
	ctx context.Context,
	token string,
	qrOrder dto.QRCodeOrder,
	date int64,
	wg *sync.WaitGroup,
	ch chan bool,
) (dto.QRCodeDataResponse, error) {
	//Block this code below until this Channel be empty (by reading with <-ch)
	ch <- true

	// TODO: rever pq precisa gerar ticket aqui. Acho que posso gerar dentro do MS de order
	// qrOrder.TicketNumber = service.orderRepository.GetNextTicketNumber(ctx, date)

	payment := dto.Payment{
		TotalPrice:  qrOrder.TotalPrice,
		PaymentType: "QR Code",
	}

	paymentResponse, err := service.paymentRepository.CreatePaymentOrder(ctx, payment)

	if err != nil {
		return dto.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	qrOrder.PaymentID = paymentResponse.PaymentId

	order := dto.Order{
		TotalPrice:   qrOrder.TotalPrice,
		CustomerCPF:  qrOrder.CustomerCPF,
		OrderProduct: []dto.OrderProduct(qrOrder.OrderProduct),
		PaymentID:    qrOrder.PaymentID,
	}

	orderResponse, err := service.orderRepository.CreatePayingOrder(ctx, order)

	if err != nil {
		return dto.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	qrCode, err := service.repository.Generate(ctx, token, order, int(orderResponse.OrderId))

	if err != nil {
		errDelete := service.orderRepository.DeleteOrder(ctx, orderResponse.OrderId)

		if errDelete != nil {
			return dto.QRCodeDataResponse{}, responses.GetResponseError(errDelete, "QRCodeGeneratorService")
		}

		return dto.QRCodeDataResponse{}, responses.GetResponseError(err, "QRCodeGeneratorService")
	}

	// Release the channel to others process be able to start a new order creation
	<-ch
	wg.Done()

	return qrCode, nil
}

func (service *FinishOrderForQRCodeUseCaseImpl) Execute(ctx context.Context, token string, form dto.ExternalPaymentEvent) error {
	if form.Topic != "merchant_order" {
		return &responses.NetworkError{
			Code: http.StatusNotAcceptable,
		}
	}

	mercadoLivrePayment, err := service.repository.GetQRCodePaymentData(ctx, token, form.Resource)

	if err != nil {
		return err
	}

	if mercadoLivrePayment.OrderStatus == "paid" {
		ids := strings.Split(mercadoLivrePayment.ExternalReference, "|")

		orderID, _ := strconv.Atoi(ids[0])
		paymentID := ids[1]

		service.paymentRepository.FinishPaymentWithSuccess(ctx, paymentID)
		service.orderRepository.FinishOrderWithPayment(ctx, uint(orderID), paymentID)
	}

	return nil
}
