package environment_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/mocks"
)

func TestEnvironment(t *testing.T) {
	t.Parallel()
	mocks.Setup()

	t.Run("got success when loading variables", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
	})

	t.Run("got success when initializing environment", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
	})

	t.Run("got success when initializing environment", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

		assert.Equal(t, "ROOT_URL", environment.GetQRCodeGatewayRootURL())
		assert.Equal(t, "token", environment.GetQRCodeGatewayToken())
		assert.Equal(t, "Region", environment.GetRegion())
		assert.Equal(t, "WEBHOOK", environment.GetWebhookMercadoLivrePaymentURL())
		assert.Equal(t, "HOST", environment.GetMongoHost())
		assert.Equal(t, "MongoDBName", environment.GetMongoDBName())
		assert.Equal(t, "OrdersRoot", environment.GetOrdersRootAPI())
	})
}
