package model

import (
	"bytes"
	"encoding/json"
	"time"
)

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
	CustomerCPF  *string        `json:"customerCPF"`
	PaymentID    string         `json:"paymentId" validate:"required"`
	OrderProduct []OrderProduct `json:"orderProducts" validate:"required"`
	TicketNumber int
}

type OrderProduct struct {
	ProductID    uint    `json:"productId" validate:"required"`
	ProductPrice float64 `json:"productPrice" validate:"required"`
}

func (o *Order) GetFormBody() (*bytes.Buffer, error) {
	jsonValue, err := json.Marshal(o)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonValue), nil
}

type OrderResponse struct {
	OrderId      uint      `json:"orderId"`
	OrderDate    time.Time `json:"orderDate"`
	TicketNumber int       `json:"ticketNumber"`
	OrderStatus  string    `json:"orderStatus"`
}

type DefaultResponse struct {
	Message string `json:"message"`
}
