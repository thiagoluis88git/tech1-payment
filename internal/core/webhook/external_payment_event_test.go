package webhook_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/webhook"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/mocks"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
)

func mockExternalPaymentEvent() dto.ExternalPaymentEvent {
	return dto.ExternalPaymentEvent{
		Resource: "Resource",
		Topic:    "Topic",
	}
}

func TestPostExternalPaymentHandler(t *testing.T) {
	t.Parallel()
	mocks.Setup()

	t.Run("got success when calling post external payment comming from Mercado Livre Webhook handler", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()

		jsonData, err := json.Marshal(mockExternalPaymentEvent())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/webhook", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		finishOrderUseCase := new(MockFinishOrderForQRCodeUseCase)

		finishOrderUseCase.On("Execute", req.Context(), "token", mockExternalPaymentEvent()).
			Return(nil)

		postExternalPaymentHandler := webhook.PostExternalPaymentEventWebhook(finishOrderUseCase)

		postExternalPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})

	t.Run("got error on FinishOrder UseCase when calling post external payment comming from Mercado Livre Webhook handler", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()

		jsonData, err := json.Marshal(mockExternalPaymentEvent())

		assert.NoError(t, err)

		body := bytes.NewBuffer(jsonData)

		req := httptest.NewRequest(http.MethodPost, "/api/webhook", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		finishOrderUseCase := new(MockFinishOrderForQRCodeUseCase)

		finishOrderUseCase.On("Execute", req.Context(), "token", mockExternalPaymentEvent()).
			Return(&responses.BusinessResponse{
				StatusCode: 500,
			})

		postExternalPaymentHandler := webhook.PostExternalPaymentEventWebhook(finishOrderUseCase)

		postExternalPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("got error on invalid json when calling post external payment comming from Mercado Livre Webhook handler", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()

		body := bytes.NewBuffer([]byte("asdff{{}"))

		req := httptest.NewRequest(http.MethodPost, "/api/webhook", body)
		req.Header.Add("Content-Type", "application/json")

		rctx := chi.NewRouteContext()

		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		recorder := httptest.NewRecorder()

		finishOrderUseCase := new(MockFinishOrderForQRCodeUseCase)

		finishOrderUseCase.On("Execute", req.Context(), "token", mockExternalPaymentEvent()).
			Return(&responses.BusinessResponse{
				StatusCode: 500,
			})

		postExternalPaymentHandler := webhook.PostExternalPaymentEventWebhook(finishOrderUseCase)

		postExternalPaymentHandler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
