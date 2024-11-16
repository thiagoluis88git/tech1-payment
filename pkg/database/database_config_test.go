package database_test

import (
	"testing"

	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/mocks"
)

func TestDatabaseConfig(t *testing.T) {
	t.Parallel()
	mocks.Setup()

	t.Run("got success when starting config database", func(t *testing.T) {
		environment.LoadEnvironmentVariables()

	})

	t.Run("got error when starting config database", func(t *testing.T) {
		environment.LoadEnvironmentVariables()
	})
}
