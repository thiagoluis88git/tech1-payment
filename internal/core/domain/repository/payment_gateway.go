package repository

import "github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"

type PaymentGateway interface {
	Pay(paymentResonse dto.PaymentResponse, payment dto.Payment) (dto.PaymentGatewayResponse, error)
}
