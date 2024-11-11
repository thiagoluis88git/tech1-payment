package handler_test

import (
	"context"
	"sync"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
)

type MockPayOrderUseCase struct {
	mock.Mock
}

type MockGetPaymentTypesUseCase struct {
	mock.Mock
}

type MockGenerateQRCodePaymentUseCase struct {
	mock.Mock
}

func (mock *MockPayOrderUseCase) Execute(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return args.Get(0).(dto.PaymentResponse), nil
}

func (m *MockGenerateQRCodePaymentUseCase) Execute(
	ctx context.Context,
	token string,
	qrOrder dto.QRCodeOrder,
	date int64,
	wg *sync.WaitGroup,
	ch chan bool,
) (dto.QRCodeDataResponse, error) {
	args := m.Called(ctx, token, qrOrder, date, wg, mock.Anything)
	err := args.Error(1)

	if err != nil {
		return dto.QRCodeDataResponse{}, err
	}

	return args.Get(0).(dto.QRCodeDataResponse), nil
}

func (mock *MockGetPaymentTypesUseCase) Execute() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}
