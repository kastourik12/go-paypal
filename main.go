package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kastouri/payment-api/paypal"
	"log"
	"os"
)

var (
	server        *gin.Engine
	ctx           context.Context
	paypalClient  *paypal.CustomClient
	paypalService *paypal.Service
	paypalHandler *paypal.Handler
)

func init() {
	godotenv.Load(".env")
	ctx = context.TODO()
	server = gin.Default()
	paypalClient = paypal.NewCustomClient(os.Getenv("PAYPAL_CLIENT_ID"), os.Getenv("PAYPAL_CLIENT_SECRET"))
	paypalService = paypal.New(*paypalClient)
	paypalHandler = paypal.NewHandler(*paypalService)
}
func main() {
	basePath := server.Group("/v1")
	paypalHandler.RegisterRoutes(basePath)
	log.Fatalf(server.Run(":8080").Error())
}
