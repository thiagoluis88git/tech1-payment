package repositories_test

import (
	"context"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"
	"github.com/thiagoluis88git/tech1-payment/pkg/database"
	"gorm.io/driver/mysql"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MockCognitoRemoteDataSource struct {
	mock.Mock
}

func (mock *MockCognitoRemoteDataSource) SignUp(user *model.Customer) error {
	args := mock.Called(user)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockCognitoRemoteDataSource) SignUpAdmin(user *model.UserAdmin) error {
	args := mock.Called(user)
	err := args.Error(0)

	if err != nil {
		return err
	}

	return nil
}

func (mock *MockCognitoRemoteDataSource) Login(cpf string) (string, error) {
	args := mock.Called(cpf)
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

func (mock *MockCognitoRemoteDataSource) LoginUnknown() (string, error) {
	args := mock.Called()
	err := args.Error(1)

	if err != nil {
		return "", err
	}

	return args.Get(0).(string), nil
}

type RepositoryTestSuite struct {
	suite.Suite
	ctx                context.Context
	db                 *database.Database
	pgContainer        *postgres.PostgresContainer
	pgConnectionString string
}

func (suite *RepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := postgres.RunContainer(
		suite.ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("notesdb"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	suite.NoError(err)

	connStr, err := pgContainer.ConnectionString(suite.ctx, "sslmode=disable")
	suite.NoError(err)

	db, err := gorm.Open(pg.Open(connStr), &gorm.Config{})
	suite.NoError(err)

	suite.pgContainer = pgContainer
	suite.pgConnectionString = connStr
	suite.db = &database.Database{Connection: db}

	sqlDB, err := suite.db.Connection.DB()
	suite.NoError(err)

	err = sqlDB.Ping()
	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	err := suite.pgContainer.Terminate(suite.ctx)
	suite.NoError(err)
}

func (suite *RepositoryTestSuite) SetupTest() {
	err := suite.db.Connection.AutoMigrate(
		&model.Product{},
		&model.ProductImage{},
		&model.ComboProduct{},
		&model.Order{},
		&model.OrderProduct{},
		&model.OrderTicketNumber{},
		&model.Customer{},
		&model.Payment{},
	)
	suite.NoError(err)
}

func (suite *RepositoryTestSuite) TearDownTest() {
	suite.db.Connection.Exec("DROP TABLE IF EXISTS customers CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS products CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS product_images CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS combo_products CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS orders CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS order_products CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS payments CASCADE;")
	suite.db.Connection.Exec("DROP TABLE IF EXISTS order_ticket_numbers CASCADE;")
}

func SetupDBMocks() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return gormDB, mock, err
}
