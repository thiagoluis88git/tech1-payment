package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thiagoluis88git/tech1-payment/internal/core/data/remote"
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/repositories"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1-payment/internal/core/handler"
	"github.com/thiagoluis88git/tech1-payment/internal/core/webhook"
	external "github.com/thiagoluis88git/tech1-payment/internal/integrations"
	integrationDS "github.com/thiagoluis88git/tech1-payment/internal/integrations/remote"
	extRepo "github.com/thiagoluis88git/tech1-payment/internal/integrations/repositories"
	"github.com/thiagoluis88git/tech1-payment/pkg/database"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/httpserver"
	"github.com/thiagoluis88git/tech1-payment/pkg/responses"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mvrilo/go-redoc"

	"github.com/go-chi/chi/v5"

	_ "github.com/thiagoluis88git/tech1-payment/docs"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Tech1 API Docs
// @version 1.0
// @description This is the API for the Tech1 Fiap Project.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localshot:3210
// @BasePath /
func main() {
	environment.LoadEnvironmentVariables()

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    *environment.RedocFolderPath,
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/docs",
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(environment.GetMongoHost()))

	if err != nil {
		panic(fmt.Sprintf("could not open database client: %v", err.Error()))
	}

	db, err := database.ConfigMongo(client, environment.GetMongoDBName())

	if err != nil {
		panic(fmt.Sprintf("could not open database: %v", err.Error()))
	}

	router := chi.NewRouter()
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)

	httpClient := httpserver.NewHTTPClient()

	orderDS := remote.NewOrderRemoteDataSource(httpClient, environment.GetOrdersRootAPI())

	paymentRepo := repositories.NewPaymentRepository(db)
	orderRepo := repositories.NewOrderRepository(orderDS)
	paymentGateway := external.NewPaymentGateway()
	payOrderUseCase := usecases.NewPayOrderUseCase(paymentRepo, paymentGateway)
	getPaymentTypesUseCase := usecases.NewGetPaymentTypesUseCasee(paymentRepo)

	qrCodeRemoteDataSource := integrationDS.NewMercadoLivreDataSource(httpClient)
	extQRCodeGeneratorRepository := extRepo.NewMercadoLivreRepository(qrCodeRemoteDataSource)
	generateQRCodePaymentUseCase := usecases.NewGenerateQRCodePaymentUseCase(
		extQRCodeGeneratorRepository,
		orderRepo,
		paymentRepo,
	)

	finishOrderForQRCodeUseCase := usecases.NewFinishOrderForQRCodeUseCase(
		extQRCodeGeneratorRepository,
		orderRepo,
		paymentRepo,
	)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		httpserver.SendResponseSuccess(w, &responses.BusinessResponse{
			StatusCode: 200,
			Message:    "ok",
		})
	})

	router.Post("/api/qrcode/generate", handler.GenerateQRCodeHandler(generateQRCodePaymentUseCase))
	router.Post("/api/webhook/ml/payment", webhook.PostExternalPaymentEventWebhook(finishOrderForQRCodeUseCase))

	router.Get("/api/payments/types", handler.GetPaymentTypeHandler(getPaymentTypesUseCase))
	router.Post("/api/payments", handler.CreatePaymentHandler(payOrderUseCase))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3210/swagger/doc.json"),
	))

	go http.ListenAndServe(":3211", doc.Handler())

	server := httpserver.New(router)
	server.Start()
}
