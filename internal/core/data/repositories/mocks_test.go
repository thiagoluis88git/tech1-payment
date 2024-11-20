package repositories_test

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
)

type MockOrderRemoteDataSource struct {
	mock.Mock
}

func (mock *MockOrderRemoteDataSource) CreatePayingOrder(ctx context.Context, order model.Order) (model.OrderResponse, error) {
	args := mock.Called(ctx, order)
	err := args.Error(1)

	if err != nil {
		return model.OrderResponse{}, err
	}

	return args.Get(0).(model.OrderResponse), nil
}

func (mock *MockOrderRemoteDataSource) DeleteOrder(ctx context.Context, orderID uint) error {
	args := mock.Called(ctx, orderID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRemoteDataSource) FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID string) error {
	args := mock.Called(ctx, orderID, paymentID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}
