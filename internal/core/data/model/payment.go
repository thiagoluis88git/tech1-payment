package model

import "time"

const (
	PaymentPayingStatus = "Pagando"
	PaymentPayedStatus  = "Pago"
	PaymentErrorStatus  = "Erro"

	PaymentCreditType = "Crédito"
	PaymentQRCodeType = "QR Code (Mercado Pago)"
)

type Payment struct {
	CustomerCPF   *string   `bson:"customerCPF"`
	TotalPrice    float64   `bson:"totalPrice"`
	PaymentStatus string    `bson:"paymentStatus"`
	PaymentType   string    `bson:"paymentType"`
	PaymentDate   time.Time `bson:"paymentDate"`
}
