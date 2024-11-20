package remote

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/pkg/httpserver"
)

type OrderRemoteDataSource interface {
	CreatePayingOrder(ctx context.Context, order model.Order) (model.OrderResponse, error)
	DeleteOrder(ctx context.Context, orderID uint) error
	FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID string) error
}

type OrderRemoteDataSourceImpl struct {
	client  *http.Client
	rootURL string
}

func NewOrderRemoteDataSource(client *http.Client, rootURL string) OrderRemoteDataSource {
	return &OrderRemoteDataSourceImpl{
		client:  client,
		rootURL: rootURL,
	}
}

func (ds *OrderRemoteDataSourceImpl) CreatePayingOrder(ctx context.Context, order model.Order) (model.OrderResponse, error) {
	endpoint := fmt.Sprintf("%v/%v", ds.rootURL, "orders")

	body, err := order.GetFormBody()

	if err != nil {
		return model.OrderResponse{}, err
	}

	response, err := httpserver.DoPostRequest(
		ctx,
		ds.client,
		endpoint,
		body,
		nil,
		model.OrderResponse{},
	)

	if err != nil {
		return model.OrderResponse{}, err
	}

	return response, nil
}

func (ds *OrderRemoteDataSourceImpl) DeleteOrder(ctx context.Context, orderID uint) error {
	endpoint := fmt.Sprintf("%v/%v/%d", ds.rootURL, "orders", orderID)

	_, err := httpserver.DoNoBodyRequest(
		ctx,
		ds.client,
		http.MethodDelete,
		endpoint,
		nil,
		model.DefaultResponse{},
	)

	if err != nil {
		return err
	}

	return nil
}

func (ds *OrderRemoteDataSourceImpl) FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID string) error {
	endpoint := fmt.Sprintf("%v/%v/%v/%d/%v", ds.rootURL, "orders", "finish", orderID, paymentID)

	_, err := httpserver.DoNoBodyRequest(
		ctx,
		ds.client,
		http.MethodPut,
		endpoint,
		nil,
		model.DefaultResponse{},
	)

	if err != nil {
		return err
	}

	return nil
}
