package dto

import "time"

type Order struct {
	OrderStatus  string
	TotalPrice   float64        `json:"totalPrice" validate:"required"`
	CustomerCPF  *string        `json:"customerCPF"`
	PaymentID    string         `json:"paymentId" validate:"required"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
}

type QRCodeOrder struct {
	OrderStatus  string
	TotalPrice   float64        `json:"totalPrice" validate:"required"`
	CustomerCPF  *string        `json:"customerCPF"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
	PaymentID    string
}

type OrderProduct struct {
	ProductID    uint    `json:"productId" validate:"required"`
	ProductPrice float64 `json:"productPrice" validate:"required"`
}

type OrderResponse struct {
	OrderId      uint      `json:"orderId"`
	OrderDate    time.Time `json:"orderDate"`
	TicketNumber int       `json:"ticketNumber"`
	OrderStatus  string    `json:"orderStatus"`
}
