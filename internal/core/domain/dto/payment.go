package dto

import "time"

type Payment struct {
	CustomerCPF *string `json:"customerCPF"`
	TotalPrice  float64 `json:"totalPrice" validate:"required"`
	PaymentType string  `json:"paymentType" validate:"required"`
}

type PaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	PaymentGatewayId string    `json:"paymentGatewayId"`
	PaymentDate      time.Time `json:"paymentDate"`
}
