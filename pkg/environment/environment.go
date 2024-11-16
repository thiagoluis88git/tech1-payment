package environment

import (
	"flag"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	RedocFolderPath *string = flag.String("PATH_REDOC_FOLDER", "/docs/swagger.json", "Swagger docs folder")

	localDev = flag.String("localDev", "false", "local development")

	singleton *Environment
)

const (
	QRCodeGatewayRootURL          = "QR_CODE_GATEWAY_ROOT_URL"
	QRCodeGatewayToken            = "QR_CODE_GATEWAY_TOKEN"
	WebhookMercadoLivrePaymentURL = "WEBHOOK_MERCADO_LIVRE_PAYMENT"
	MongoHost                     = "MONGO_HOST"
	MongoPassword                 = "MONGO_PASSWORD"
	MongoDBName                   = "MONGO_DB_NAME"
	Region                        = "AWS_REGION"
	OrdersRootAPI                 = "ORDERS_ROOT_API"
	passwordFieldToReplace        = "<db_password>"
)

type Environment struct {
	qrCodeGatewayRootURL          string
	qrCodeGatewayToken            string
	webhookMercadoLivrePaymentURL string
	mongoHost                     string
	mongoPassword                 string
	mongoDBName                   string
	region                        string
	ordersRootAPI                 string
}

func LoadEnvironmentVariables() {
	flag.Parse()

	if localFlag := *localDev; localFlag != "false" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err.Error())
		}
	}

	qrCodeGatewayRootURL := getEnvironmentVariable(QRCodeGatewayRootURL)
	qrCodeGatewayToken := getEnvironmentVariable(QRCodeGatewayToken)
	webhookMercadoLivrePaymentURL := getEnvironmentVariable(WebhookMercadoLivrePaymentURL)
	mongoHost := getEnvironmentVariable(MongoHost)
	mongoPassword := getEnvironmentVariable(MongoPassword)
	mongoDBName := getEnvironmentVariable(MongoDBName)
	region := getEnvironmentVariable(Region)
	ordersRootAPI := getEnvironmentVariable(OrdersRootAPI)

	once := &sync.Once{}

	once.Do(func() {
		singleton = &Environment{
			qrCodeGatewayRootURL:          qrCodeGatewayRootURL,
			qrCodeGatewayToken:            qrCodeGatewayToken,
			webhookMercadoLivrePaymentURL: webhookMercadoLivrePaymentURL,
			mongoHost:                     mongoHost,
			mongoPassword:                 mongoPassword,
			mongoDBName:                   mongoDBName,
			region:                        region,
			ordersRootAPI:                 ordersRootAPI,
		}
	})
}

func getEnvironmentVariable(key string) string {
	value, hashKey := os.LookupEnv(key)

	if !hashKey {
		log.Fatalf("There is no %v environment variable", key)
	}

	return value
}

func GetWebhookMercadoLivrePaymentURL() string {
	return singleton.webhookMercadoLivrePaymentURL
}

func GetQRCodeGatewayRootURL() string {
	return singleton.qrCodeGatewayRootURL
}

func GetQRCodeGatewayToken() string {
	return singleton.qrCodeGatewayToken
}

func GetMongoHost() string {
	host := singleton.mongoHost
	host = strings.ReplaceAll(host, passwordFieldToReplace, url.QueryEscape(singleton.mongoPassword))

	return host
}

func GetMongoDBName() string {
	return singleton.mongoDBName
}

func GetRegion() string {
	return singleton.region
}

func GetOrdersRootAPI() string {
	return singleton.ordersRootAPI
}
