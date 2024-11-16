package repositories

import (
	"context"

	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/repository"
	"github.com/thiagoluis88git/tech1-payment/pkg/database"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
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
	}

	result, err := repository.db.Conn.Collection(paymentCollectionName).InsertOne(ctx, paymentEntity)

	if err != nil {
		return dto.PaymentResponse{}, responses.GetDatabaseError(err)
	}

	return dto.PaymentResponse{
		PaymentId: result.InsertedID.(string),
	}, nil
}

func (repository *PaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId string) error {
	paymentEntity := model.Payment{
		PaymentStatus: model.PaymentErrorStatus,
	}

	_, err := repository.db.Conn.Collection(paymentCollectionName).UpdateByID(ctx, paymentId, paymentEntity)

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}

func (repository *PaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId string) error {
	paymentEntity := model.Payment{
		PaymentStatus: model.PaymentPayedStatus,
	}

	_, err := repository.db.Conn.Collection(paymentCollectionName).UpdateByID(ctx, paymentId, paymentEntity)

	if err != nil {
		return responses.GetDatabaseError(err)
	}

	return nil
}
