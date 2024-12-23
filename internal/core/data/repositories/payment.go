package repositories

import (
	"context"
	"time"

	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1-payment/pkg/database"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	paymentCollectionName = "payments"
)

type PaymentRepository struct {
	db *database.Database
}

func NewPaymentRepository(db *database.Database) repository.PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (repository *PaymentRepository) GetPaymentTypes() []string {
	return []string{
		model.PaymentQRCodeType,
		model.PaymentCreditType,
	}
}

func (repository *PaymentRepository) CreatePaymentOrder(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	paymentEntity := model.Payment{
		CustomerCPF:   payment.CustomerCPF,
		TotalPrice:    payment.TotalPrice,
		PaymentType:   payment.PaymentType,
		PaymentStatus: model.PaymentPayingStatus,
		PaymentDate:   time.Now(),
	}

	result, err := repository.db.Conn.
		Collection(paymentCollectionName).
		InsertOne(ctx, paymentEntity)

	if err != nil {
		return dto.PaymentResponse{}, responses.GetDatabaseError(err)
	}

	objID := result.InsertedID.(primitive.ObjectID)

	return dto.PaymentResponse{
		PaymentId: objID.Hex(),
	}, nil
}

func (repository *PaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId string) error {
	return repository.updateWithStatus(ctx, paymentId, model.PaymentErrorStatus)
}

func (repository *PaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId string) error {
	return repository.updateWithStatus(ctx, paymentId, model.PaymentPayedStatus)
}

func (repository *PaymentRepository) updateWithStatus(ctx context.Context, paymentID string, status string) error {
	update := bson.D{{
		Key: "$set", Value: bson.D{{Key: "paymentStatus", Value: status}},
	}}

	id, err := primitive.ObjectIDFromHex(paymentID)

	if err != nil {
		return err
	}

	_, err = repository.db.Conn.
		Collection(paymentCollectionName).
		UpdateByID(ctx, id, update)

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}
