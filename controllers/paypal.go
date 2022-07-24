package controllers

import (
	"github.com/gin-gonic/gin"
	"kastouri/payment-api/requests"
	"kastouri/payment-api/services"
)

type Paypal struct {
	service *services.Paypal
}

func NewPaypalController(s *services.Paypal) *Paypal {
	return &Paypal{
		service: s,
	}
}
func (c *Paypal) CreatePayment(ctx *gin.Context) {
	var paypalRequest requests.Paypal
	ctx.BindJSON(&paypalRequest)
	successLink, err := c.service.CreatePayment(paypalRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, successLink)
}

func (c *Paypal) ExecutePayment(ctx *gin.Context) {
	payerID := ctx.Query("PayerID")
	paymentID := ctx.Query("paymentId")
	err := c.service.ExecutePayment(paymentID, payerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func (c *Paypal) RegisterRoutes(rg *gin.RouterGroup) {
	paypalRoute := rg.Group("/paypal")
	paypalRoute.POST("/create", c.CreatePayment)
	paypalRoute.GET("/execute", c.ExecutePayment)
}
