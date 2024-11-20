package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
)

func TestOrdersRepository(t *testing.T) {
	t.Parallel()

	t.Run("got success when creating paying order repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("CreatePayingOrder", ctx, model.Order{
			TotalPrice: 234.45,
			OrderProduct: []model.OrderProduct{
				{
					ProductID: uint(9),
				},
			},
		}).Return(model.OrderResponse{
			OrderId: uint(34),
		}, nil)

		response, err := sut.CreatePayingOrder(ctx, dto.Order{
			TotalPrice: 234.45,
			OrderProduct: []dto.OrderProduct{
				{
					ProductID: uint(9),
				},
			},
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)

		assert.Equal(t, uint(34), response.OrderId)
	})

	t.Run("got error when creating paying order repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("CreatePayingOrder", ctx, model.Order{
			TotalPrice: 234.45,
			OrderProduct: []model.OrderProduct{
				{
					ProductID: uint(9),
				},
			},
		}).Return(model.OrderResponse{}, &responses.NetworkError{
			Code: 500,
		})

		response, err := sut.CreatePayingOrder(ctx, dto.Order{
			TotalPrice: 234.45,
			OrderProduct: []dto.OrderProduct{
				{
					ProductID: uint(9),
				},
			},
		})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when deleting order repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("DeleteOrder", ctx, uint(3)).Return(nil)

		err := sut.DeleteOrder(ctx, uint(3))

		assert.NoError(t, err)
	})

	t.Run("got error when deleting order repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("DeleteOrder", ctx, uint(3)).Return(&responses.NetworkError{
			Code: 500,
		})

		err := sut.DeleteOrder(ctx, uint(3))

		assert.Error(t, err)
	})

	t.Run("got success when finishing order with payment repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("FinishOrderWithPayment", ctx, uint(3), "paymentID").Return(nil)

		err := sut.FinishOrderWithPayment(ctx, uint(3), "paymentID")

		assert.NoError(t, err)
	})

	t.Run("got error when finishing order with payment repository", func(t *testing.T) {
		ds := new(MockOrderRemoteDataSource)
		sut := repositories.NewOrderRepository(ds)

		ctx := context.TODO()

		ds.On("FinishOrderWithPayment", ctx, uint(3), "paymentID").Return(&responses.NetworkError{
			Code: 500,
		})

		err := sut.FinishOrderWithPayment(ctx, uint(3), "paymentID")

		assert.Error(t, err)
	})
}
