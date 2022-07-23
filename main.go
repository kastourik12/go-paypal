package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/netlify/PayPal-Go-SDK"
	"kastouri/payment-api/controllers"
	"kastouri/payment-api/services"
	"log"
	"os"
)

var (
	server           *gin.Engine
	ctx              context.Context
	client           *paypalsdk.Client
	paypalController *controllers.Paypal
	paypalService    *services.Paypal
	logger           *log.Logger
)

func init() {
	logger = log.Default()
	godotenv.Load(".env")
	ctx = context.TODO()
	server = gin.Default()
	client, err := paypalsdk.NewClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_CLIENT_SECRET"), paypalsdk.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	accessToken, err := client.GetAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(accessToken.Token)
	paypalService = services.NewPaypalService(client, logger)
	paypalController = controllers.NewPaypalController(paypalService)

}
func main() {
	basePath := server.Group("/v1")
	paypalController.RegisterRoutes(basePath)
	log.Fatalf(server.Run(":8080").Error())
}
