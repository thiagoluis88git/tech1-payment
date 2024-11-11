package repository

import (
	"context"

	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
)

type OrderRepository interface {
	CreatePayingOrder(ctx context.Context, order dto.Order) (dto.OrderResponse, error)
	DeleteOrder(ctx context.Context, orderID uint) error
	FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID uint) error
}
