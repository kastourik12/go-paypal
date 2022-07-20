package paypal

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: &service,
	}
}
func (h *Handler) CreatePayment(ctx *gin.Context) {

	var paymentRequest PaymentRequest
	err := ctx.BindJSON(&paymentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	err = h.Service.CreatePayment(paymentRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, paymentRequest)

}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	pRoute := r.Group("/paypal")
	pRoute.POST("/create", h.CreatePayment)

}
