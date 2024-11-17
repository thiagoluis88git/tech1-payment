package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/pkg/database"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestPaymentRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("got success when creating payment order respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response, err := repo.CreatePaymentOrder(context.TODO(), dto.Payment{})
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	mt.Run("got error when creating payment order respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		response, err := repo.CreatePaymentOrder(context.TODO(), dto.Payment{})
		assert.Error(t, err)
		assert.Empty(t, response)
	})

	mt.Run("got success when updating payment order with PAYED respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err = repo.FinishPaymentWithSuccess(context.TODO(), "6738ccbbf05e936dcc11a986")
		assert.NoError(t, err)
	})

	mt.Run("got error on invalid object id hex when updating payment order with PAYED respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err = repo.FinishPaymentWithSuccess(context.TODO(), "123")
		assert.Error(t, err)
	})

	mt.Run("got error when updating payment order with PAYED respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err = repo.FinishPaymentWithSuccess(context.TODO(), "6738ccbbf05e936dcc11a986")
		assert.Error(t, err)
	})

	mt.Run("got success when updating payment order with ERROR respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err = repo.FinishPaymentWithError(context.TODO(), "6738ccbbf05e936dcc11a986")
		assert.NoError(t, err)
	})

	mt.Run("got error on invalid object id hex when updating payment order with ERROR respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err = repo.FinishPaymentWithError(context.TODO(), "123")
		assert.Error(t, err)
	})

	mt.Run("got error when updating payment order with ERROR respository", func(mt *mtest.T) { // test code
		database, err := database.ConfigMongo(mt.Client, mt.Name())

		assert.NoError(t, err)

		repo := repositories.NewPaymentRepository(database)
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "duplicate key error",
		}))

		err = repo.FinishPaymentWithError(context.TODO(), "6738ccbbf05e936dcc11a986")
		assert.Error(t, err)
	})
}
