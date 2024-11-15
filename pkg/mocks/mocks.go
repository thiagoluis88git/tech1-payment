package mocks

import (
	"net/http"
	"os"

	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
)

func Setup() {
	os.Setenv(environment.QRCodeGatewayRootURL, "ROOT_URL")
	os.Setenv(environment.DBHost, "HOST")
	os.Setenv(environment.DBPort, "1234")
	os.Setenv(environment.DBUser, "User")
	os.Setenv(environment.DBPassword, "Pass")
	os.Setenv(environment.DBName, "Name")
	os.Setenv(environment.WebhookMercadoLivrePaymentURL, "WEBHOOK")
	os.Setenv(environment.QRCodeGatewayToken, "token")
	os.Setenv(environment.Region, "Region")
	os.Setenv(environment.OrdersRootAPI, "OrdersRoot")
}

type MockRoundTripper struct {
	Response *http.Response
}

func (trip *MockRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return trip.Response, nil
}
