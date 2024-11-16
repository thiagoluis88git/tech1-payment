package environment_test

import (
	"testing"

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
}
