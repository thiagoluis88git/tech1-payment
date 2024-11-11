package usecases

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
)

var (
	paymentCreation = dto.Payment{
		TotalPrice:  1234,
		PaymentType: "Crédito",
	}

	paymentResponse = dto.PaymentResponse{
		PaymentId:        1,
		PaymentGatewayId: "123",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}

	paymentGatewayResponse = dto.PaymentGatewayResponse{
		PaymentGatewayId: "1234",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}
)

type MockOrderRepository struct {
	mock.Mock
}

type MockCustomerRepository struct {
	mock.Mock
}

type MockPaymentRepository struct {
	mock.Mock
}

type MockPaymentGatewayRepository struct {
	mock.Mock
}

type MockProductRepository struct {
	mock.Mock
}

type MockUserAdminRepository struct {
	mock.Mock
}

type MockQRCodePaymentRepository struct {
	mock.Mock
}

func (mock *MockOrderRepository) CreatePayingOrder(ctx context.Context, order dto.Order) (dto.OrderResponse, error) {
	args := mock.Called(ctx, order)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockOrderRepository) DeleteOrder(ctx context.Context, orderID uint) error {
	args := mock.Called(ctx, orderID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID uint) error {
	args := mock.Called(ctx, orderID, paymentID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) GetOrderById(ctx context.Context, orderId uint) (dto.OrderResponse, error) {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersToPrepare(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersToFollow(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersWaitingPayment(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) UpdateToPreparing(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToDone(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) GetNextTicketNumber(ctx context.Context, date int64) int {
	args := mock.Called(ctx, date)
	return args.Get(0).(int)
}

func (mock *MockPaymentRepository) GetPaymentTypes() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

func (mock *MockPaymentRepository) CreatePaymentOrder(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return args.Get(0).(dto.PaymentResponse), nil
}

func (mock *MockPaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentGatewayRepository) Pay(paymentResonse dto.PaymentResponse, payment dto.Payment) (dto.PaymentGatewayResponse, error) {
	args := mock.Called(paymentResonse, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentGatewayResponse{}, err
	}

	return args.Get(0).(dto.PaymentGatewayResponse), nil
}

func (mock *MockQRCodePaymentRepository) Generate(
	ctx context.Context,
	token string,
	form dto.Order,
	orderID int,
) (dto.QRCodeDataResponse, error) {
	args := mock.Called(ctx, token, form, orderID)
	err := args.Error(1)

	if err != nil {
		return dto.QRCodeDataResponse{}, err
	}

	return args.Get(0).(dto.QRCodeDataResponse), nil
}

func (mock *MockQRCodePaymentRepository) GetQRCodePaymentData(
	ctx context.Context,
	token string,
	endpoint string,
) (dto.ExternalPaymentInformation, error) {
	args := mock.Called(ctx, token, endpoint)
	err := args.Error(1)

	if err != nil {
		return dto.ExternalPaymentInformation{}, err
	}

	return args.Get(0).(dto.ExternalPaymentInformation), nil
}
