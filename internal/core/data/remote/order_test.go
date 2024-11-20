package remote_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/remote"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/mocks"
)

func TestOrdersRemote(t *testing.T) {
	t.Parallel()
	mocks.Setup()

	t.Run("got success when creating payment order remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString(MockOrder)

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		response, err := ds.CreatePayingOrder(context.TODO(), model.Order{})

		assert.NoError(t, err)
		assert.Equal(t, 3, response.TicketNumber)
	})

	t.Run("got error on invalid json when creating payment order remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString("sdasd{{}")

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		response, err := ds.CreatePayingOrder(context.TODO(), model.Order{})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when deleting order remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString(MockDefaultOrder)

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		err = ds.DeleteOrder(context.TODO(), uint(2))

		assert.NoError(t, err)
	})

	t.Run("got error when deleting order remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString("fff}}")

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		err = ds.DeleteOrder(context.TODO(), uint(2))

		assert.Error(t, err)
	})

	t.Run("got success when finishing order with payment remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString(MockDefaultOrder)

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		err = ds.FinishOrderWithPayment(context.TODO(), uint(2), "paymentID")

		assert.NoError(t, err)
	})

	t.Run("got error when finishing order with payment remote", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		recorder := httptest.NewRecorder()
		recorder.Header().Add("Content-Type", "application/json")
		recorder.WriteHeader(http.StatusOK)
		_, err := recorder.WriteString("sdf{{}")

		assert.NoError(t, err)

		resultExpected := recorder.Result()

		mockClient := &http.Client{
			Transport: &mocks.MockRoundTripper{
				Response: resultExpected,
			},
		}

		ds := remote.NewOrderRemoteDataSource(mockClient, "OrdersRoot")

		err = ds.FinishOrderWithPayment(context.TODO(), uint(2), "paymentID")

		assert.Error(t, err)
	})
}
