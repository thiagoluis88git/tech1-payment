package model

import "time"

const (
	OrderStatusPaying       = "Em pagamento"
	OrderStatusCreated      = "Criado"
	OrderStatusPreparing    = "Preparando"
	OrderStatusDone         = "Finalizado"
	OrderStatusDelivered    = "Entregue"
	OrderStatusNotDelivered = "Não entregue"
)

type Order struct {
	OrderStatus  string
	TotalPrice   float64        `json:"totalPrice" validate:"required"`
	CustomerID   *uint          `json:"customerId"`
	PaymentID    uint           `json:"paymentId" validate:"required"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
	TicketNumber int
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
