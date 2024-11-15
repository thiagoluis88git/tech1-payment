package environment_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
)

func setup() {
	os.Setenv(environment.DBHost, "DBHost")
	os.Setenv(environment.DBPassword, "DBPassword")
	os.Setenv(environment.DBName, "DBName")
	os.Setenv(environment.DBPort, "DBPort")
	os.Setenv(environment.DBUser, "DBUser")
	os.Setenv(environment.QRCodeGatewayRootURL, "QRCodeGatewayRootURL")
	os.Setenv(environment.QRCodeGatewayToken, "QRCodeGatewayToken")
	os.Setenv(environment.Region, "Region")
	os.Setenv(environment.WebhookMercadoLivrePaymentURL, "WebhookMercadoLivrePaymentURL")
}

func TestEnvironment(t *testing.T) {
	t.Parallel()
	setup()

	t.Run("got success when loading variables", func(t *testing.T) {
		t.Parallel()

		environment.LoadEnvironmentVariables()
	})

	t.Run("got success when initializing environment", func(t *testing.T) {
		t.Parallel()
		environment.LoadEnvironmentVariables()

		assert.Equal(t, "DBHost", environment.GetDBHost())
		assert.Equal(t, "DBPassword", environment.GetDBPassword())
		assert.Equal(t, "DBPort", environment.GetDBPort())
		assert.Equal(t, "DBName", environment.GetDBName())
		assert.Equal(t, "DBUser", environment.GetDBUser())
		assert.Equal(t, "QRCodeGatewayRootURL", environment.GetQRCodeGatewayRootURL())
		assert.Equal(t, "QRCodeGatewayToken", environment.GetQRCodeGatewayToken())
		assert.Equal(t, "Region", environment.GetRegion())
		assert.Equal(t, "WebhookMercadoLivrePaymentURL", environment.GetWebhookMercadoLivrePaymentURL())
	})
}
