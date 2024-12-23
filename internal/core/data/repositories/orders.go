package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/remote"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/repository"
)

type OrderRepositoryImpl struct {
	ds remote.OrderRemoteDataSource
}

func NewOrderRepository(ds remote.OrderRemoteDataSource) repository.OrderRepository {
	return OrderRepositoryImpl{
		ds: ds,
	}
}

func (repo OrderRepositoryImpl) CreatePayingOrder(ctx context.Context, order dto.Order) (dto.OrderResponse, error) {
	products := make([]model.OrderProduct, 0)

	for _, value := range order.OrderProduct {
		products = append(products, model.OrderProduct(value))
	}

	input := model.Order{
		TotalPrice:   order.TotalPrice,
		OrderStatus:  order.OrderStatus,
		CustomerCPF:  order.CustomerCPF,
		PaymentID:    order.PaymentID,
		OrderProduct: products,
	}

	response, err := repo.ds.CreatePayingOrder(ctx, input)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return dto.OrderResponse{
		OrderId:      response.OrderId,
		OrderDate:    response.OrderDate,
		TicketNumber: response.TicketNumber,
		OrderStatus:  response.OrderStatus,
	}, nil
}

func (repo OrderRepositoryImpl) DeleteOrder(ctx context.Context, orderID uint) error {
	err := repo.ds.DeleteOrder(ctx, orderID)

	if err != nil {
		return err
	}

	return nil
}

func (repo OrderRepositoryImpl) FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID string) error {
	err := repo.ds.FinishOrderWithPayment(ctx, orderID, paymentID)

	if err != nil {
		return err
	}

	return nil
}
