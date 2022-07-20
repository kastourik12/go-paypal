package paypal

import (
	"fmt"
	"github.com/logpacker/PayPal-Go-SDK"
	"log"
	"time"
)

type CustomClient struct {
	ClientId     string
	ClientSecret string
	token        string
	expiresIn    time.Time
	client       *paypal.Client
	refreshToken string
}

func NewCustomClient(clientId, clientSecret string) *CustomClient {
	client, err := paypal.NewClient(clientId, clientSecret, paypal.APIBaseLive)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	token, err := client.GetAccessToken()
	if err != nil {
		token.Token = ""
		token.RefreshToken = ""
		token.ExpiresIn = 0
	}

	return &CustomClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		token:        token.Token,
		refreshToken: token.RefreshToken,
		expiresIn:    time.Now().Add(time.Duration(token.ExpiresIn) * time.Second),
		client:       client,
	}
}
func (c *CustomClient) GetToken() string {
	return c.token
}
func (c *CustomClient) RefreshToken() error {
	if c.expiresIn.After(time.Now()) {
		token, err := c.client.GetAccessToken()
		if err != nil {
			return err
		}
		c.token = token.Token
		c.refreshToken = token.RefreshToken
		c.expiresIn = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
		return nil
	}
	return nil
}
func (c *CustomClient) GetClient() string { return c.client.Token.Token }

func (c *CustomClient) CreatePayment(request PaymentRequest) error {
	amount := new(paypal.PurchaseUnitAmount)
	amount.Currency = request.currency
	amount.Value = request.amount
	payment := new(paypal.PurchaseUnitRequest)
	payment.Amount = amount
	payer := new(paypal.CreateOrderPayer)
	var payments []paypal.PurchaseUnitRequest
	payments = append(payments, *payment)
	appContext := new(paypal.ApplicationContext)
	appContext.BrandName = "Company Name"
	appContext.LandingPage = "BILLING"
	appContext.ReturnURL = "http://localhost:9090/excute"
	appContext.CancelURL = "http://localhost:3000/cancel"
	order, err := c.client.CreateOrder("sale", payments, payer, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(order.ID)
	return nil
}
func (c *CustomClient) ExecutePayment(orderId string) error {
	_, err := c.client.CaptureOrder(orderId, paypal.CaptureOrderRequest{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
