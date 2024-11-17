package httpserver_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/pkg/httpserver"
)

var (
	token = "token"
)

func TestHttpClient(t *testing.T) {
	t.Parallel()

	t.Run("got success when creating http client", func(t *testing.T) {
		t.Parallel()

		client := httpserver.NewHTTPClient()

		assert.NotEmpty(t, client)
	})

	t.Run("got success when calling DoNoBodyRequest client", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpserver.SendResponseSuccess(w, dto.Token{
				AccessToken: "TOKEN",
			})
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		response, err := httpserver.DoNoBodyRequest(context.TODO(), ts.Client(), http.MethodGet, ts.URL, nil, dto.Token{})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error on invalid json when calling DoNoBodyRequest client", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		response, err := httpserver.DoNoBodyRequest(context.TODO(), ts.Client(), http.MethodGet, ts.URL, nil, dto.Token{})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error when calling DoNoBodyRequest client", func(t *testing.T) {
		t.Parallel()

		client := httpserver.NewHTTPClient()

		assert.NotEmpty(t, client)

		response, err := httpserver.DoNoBodyRequest(context.TODO(), client, http.MethodGet, "http://localhost", nil, dto.Token{})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got success when calling DoPostRequest client", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpserver.SendResponseSuccess(w, dto.Token{
				AccessToken: "TOKEN",
			})
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		response, err := httpserver.DoPostRequest(context.TODO(), ts.Client(), ts.URL, nil, &token, dto.Token{})

		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("got error on invalid json when calling DoPostRequest client", func(t *testing.T) {
		t.Parallel()

		responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		ts := httptest.NewServer(responseHandler)
		defer ts.Close()

		response, err := httpserver.DoPostRequest(context.TODO(), ts.Client(), ts.URL, nil, &token, dto.Token{})

		assert.Error(t, err)
		assert.Empty(t, response)
	})

	t.Run("got error when calling DoPostRequest client", func(t *testing.T) {
		t.Parallel()

		client := httpserver.NewHTTPClient()

		assert.NotEmpty(t, client)

		response, err := httpserver.DoPostRequest(context.TODO(), client, "http://localhost", nil, &token, dto.Token{})

		assert.Error(t, err)
		assert.Empty(t, response)
	})
}
