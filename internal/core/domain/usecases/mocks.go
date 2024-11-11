package usecases

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
)

var (
	orderCreation = dto.Order{
		TotalPrice:   12345,
		PaymentID:    uint(1),
		TicketNumber: 1,
		OrderProduct: []dto.OrderProduct{
			{
				ProductID: 1,
			},
			{
				ProductID: 2,
			},
		},
	}
	customerId = uint(1)

	orderCreationWithCustomer = dto.Order{
		TotalPrice:   12345,
		PaymentID:    uint(1),
		TicketNumber: 1,
		CustomerID:   &customerId,
		OrderProduct: []dto.OrderProduct{
			{
				ProductID: 1,
			},
			{
				ProductID: 2,
			},
		},
	}

	customerName    = "Customer Name"
	paymentCreation = dto.Payment{
		TotalPrice:  1234,
		PaymentType: "Crédito",
	}

	paymentResponse = dto.PaymentResponse{
		PaymentId:        1,
		PaymentGatewayId: "123",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}

	paymentGatewayResponse = dto.PaymentGatewayResponse{
		PaymentGatewayId: "1234",
		PaymentDate:      time.Date(2024, 10, 10, 0, 0, 0, 0, time.Local),
	}
)

type MockOrderRepository struct {
	mock.Mock
}

type MockCustomerRepository struct {
	mock.Mock
}

type MockPaymentRepository struct {
	mock.Mock
}

type MockPaymentGatewayRepository struct {
	mock.Mock
}

type MockProductRepository struct {
	mock.Mock
}

type MockUserAdminRepository struct {
	mock.Mock
}

type MockQRCodePaymentRepository struct {
	mock.Mock
}

func (mock *MockCustomerRepository) CreateCustomer(ctx context.Context, customer dto.Customer) (uint, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return 0, err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockCustomerRepository) Login(ctx context.Context, cpf string) (string, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockCustomerRepository) LoginUnknown() (string, error) {
	args := mock.Called()
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockCustomerRepository) UpdateCustomer(ctx context.Context, customer dto.Customer) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockCustomerRepository) GetCustomerByCPF(ctx context.Context, cpf string) (dto.Customer, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockCustomerRepository) GetCustomerById(ctx context.Context, id uint) (dto.Customer, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.Customer{}, err
	}

	return args.Get(0).(dto.Customer), nil
}

func (mock *MockOrderRepository) CreateOrder(ctx context.Context, order dto.Order) (dto.OrderResponse, error) {
	args := mock.Called(ctx, order)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockOrderRepository) CreatePayingOrder(ctx context.Context, order dto.Order) (dto.OrderResponse, error) {
	args := mock.Called(ctx, order)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockOrderRepository) DeleteOrder(ctx context.Context, orderID uint) error {
	args := mock.Called(ctx, orderID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) FinishOrderWithPayment(ctx context.Context, orderID uint, paymentID uint) error {
	args := mock.Called(ctx, orderID, paymentID)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) GetOrderById(ctx context.Context, orderId uint) (dto.OrderResponse, error) {
	args := mock.Called(ctx, orderId)
	err := args.Error(1)

	if err != nil {
		return dto.OrderResponse{}, err
	}

	return args.Get(0).(dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersToPrepare(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersToFollow(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) GetOrdersWaitingPayment(ctx context.Context) ([]dto.OrderResponse, error) {
	args := mock.Called(ctx)
	err := args.Error(1)

	if err != nil {
		return []dto.OrderResponse{}, err
	}

	return args.Get(0).([]dto.OrderResponse), nil
}

func (mock *MockOrderRepository) UpdateToPreparing(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToDone(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToDelivered(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) UpdateToNotDelivered(ctx context.Context, orderId uint) error {
	args := mock.Called(ctx, orderId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockOrderRepository) GetNextTicketNumber(ctx context.Context, date int64) int {
	args := mock.Called(ctx, date)
	return args.Get(0).(int)
}

func (mock *MockPaymentRepository) GetPaymentTypes() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

func (mock *MockPaymentRepository) CreatePaymentOrder(ctx context.Context, payment dto.Payment) (dto.PaymentResponse, error) {
	args := mock.Called(ctx, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentResponse{}, err
	}

	return args.Get(0).(dto.PaymentResponse), nil
}

func (mock *MockPaymentRepository) FinishPaymentWithSuccess(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentRepository) FinishPaymentWithError(ctx context.Context, paymentId uint) error {
	args := mock.Called(ctx, paymentId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockPaymentGatewayRepository) Pay(paymentResonse dto.PaymentResponse, payment dto.Payment) (dto.PaymentGatewayResponse, error) {
	args := mock.Called(paymentResonse, payment)
	err := args.Error(1)

	if err != nil {
		return dto.PaymentGatewayResponse{}, err
	}

	return args.Get(0).(dto.PaymentGatewayResponse), nil
}

func (mock *MockProductRepository) CreateProduct(ctx context.Context, product dto.ProductForm) (uint, error) {
	args := mock.Called(ctx, product)
	err := args.Error(1)

	if err != nil {
		return 0, err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockProductRepository) GetProductsByCategory(ctx context.Context, category string) ([]dto.ProductResponse, error) {
	args := mock.Called(ctx, category)
	err := args.Error(1)

	if err != nil {
		return []dto.ProductResponse{}, err
	}

	return args.Get(0).([]dto.ProductResponse), nil
}

func (mock *MockProductRepository) GetProductById(ctx context.Context, id uint) (dto.ProductResponse, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.ProductResponse{}, err
	}

	return args.Get(0).(dto.ProductResponse), nil
}

func (mock *MockProductRepository) DeleteProduct(ctx context.Context, productId uint) error {
	args := mock.Called(ctx, productId)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockProductRepository) UpdateProduct(ctx context.Context, product dto.ProductForm) error {
	args := mock.Called(ctx, product)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockProductRepository) GetCategories() []string {
	args := mock.Called()
	return args.Get(0).([]string)
}

func (mock *MockUserAdminRepository) CreateUser(ctx context.Context, customer dto.UserAdmin) (uint, error) {
	args := mock.Called(ctx, customer)
	err := args.Error(1)

	if err != nil {
		return uint(0), err
	}

	return args.Get(0).(uint), nil
}

func (mock *MockUserAdminRepository) GetUserById(ctx context.Context, id uint) (dto.UserAdmin, error) {
	args := mock.Called(ctx, id)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockUserAdminRepository) GetUserByCPF(ctx context.Context, cpf string) (dto.UserAdmin, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return dto.UserAdmin{}, err
	}

	return args.Get(0).(dto.UserAdmin), nil
}

func (mock *MockUserAdminRepository) Login(ctx context.Context, cpf string) (string, error) {
	args := mock.Called(ctx, cpf)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockUserAdminRepository) UpdateUser(ctx context.Context, customer dto.UserAdmin) error {
	args := mock.Called(ctx, customer)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockQRCodePaymentRepository) Generate(
	ctx context.Context,
	token string,
	form dto.Order,
	orderID int,
) (dto.QRCodeDataResponse, error) {
	args := mock.Called(ctx, token, form, orderID)
	err := args.Error(1)

	if err != nil {
		return dto.QRCodeDataResponse{}, err
	}

	return args.Get(0).(dto.QRCodeDataResponse), nil
}

func (mock *MockQRCodePaymentRepository) GetQRCodePaymentData(
	ctx context.Context,
	token string,
	endpoint string,
) (dto.ExternalPaymentInformation, error) {
	args := mock.Called(ctx, token, endpoint)
	err := args.Error(1)

	if err != nil {
		return dto.ExternalPaymentInformation{}, err
	}

	return args.Get(0).(dto.ExternalPaymentInformation), nil
}
