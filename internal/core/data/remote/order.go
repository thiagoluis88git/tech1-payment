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
