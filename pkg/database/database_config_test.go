package database_test

import (
	"os"
	"testing"

	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
)

func setup() {
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

func TestDatabaseConfig(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when starting config database", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

	})

	t.Run("got error when starting config database", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

	})
}
