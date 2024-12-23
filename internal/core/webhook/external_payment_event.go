package webhook

import (
	"log"
	"net/http"

	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/dto"
	"github.com/thiagoluis88git/tech1-payment/internal/core/domain/usecases"
	"github.com/thiagoluis88git/tech1-payment/pkg/environment"
	"github.com/thiagoluis88git/tech1-payment/pkg/httpserver"
)

// @Summary Payment Webhook
// @Description Payment Webhook. This endpoint will be called when the user pays
// @Description the QRCode generated by /api/qrcode/generate [post]
// @Tags Webhook
// @Accept json
// @Produce json
// @Param externalPaymentEvent body dto.ExternalPaymentEvent true "externalPaymentEvent"
// @Success 204
// @Failure 406 "StatusNotAcceptable - Topic is not 'merchant_order'"
// @Router /api/webhook/ml/payment [post]
func PostExternalPaymentEventWebhook(finishOrderForQRCode usecases.FinishOrderForQRCodeUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var form dto.ExternalPaymentEvent

		err := httpserver.DecodeJSONBody(w, r, &form)

		if err != nil {
			log.Print("decoding mercado livre webhook body", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendBadRequestError(w, err)
			return
		}

		token := environment.GetQRCodeGatewayToken()
		err = finishOrderForQRCode.Execute(r.Context(), token, form)

		if err != nil {
			log.Print("post mercado livre webhook", map[string]interface{}{
				"error":  err.Error(),
				"status": httpserver.GetStatusCodeFromError(err),
			})
			httpserver.SendResponseError(w, err)
			return
		}

		httpserver.SendResponseNoContentSuccess(w)
	}
}
