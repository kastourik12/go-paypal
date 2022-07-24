package services

import (
	paypalsdk "github.com/netlify/PayPal-Go-SDK"
	"kastouri/payment-api/requests"
	"log"
)

type Paypal struct {
	client *paypalsdk.Client
	logger *log.Logger
}

func NewPaypalService(c *paypalsdk.Client, l *log.Logger) *Paypal {
	return &Paypal{
		client: c,
		logger: l,
	}
}
func (s *Paypal) CreatePayment(request requests.Paypal) (string, error) {
	amount := paypalsdk.Amount{
		Total:    request.Total,
		Currency: request.Currency,
	}
	payment, err := s.client.CreateDirectPaypalPayment(amount, "http://localhost:8080/v1/paypal/execute", "http://exemple.com/cancel", "")
	if err != nil {
		return "", err
	}
	var successLink string
	for _, link := range payment.Links {
		if link.Rel == "approval_url" {
			successLink = link.Href

		}
	}
	return successLink, nil
}
func (s *Paypal) ExecutePayment(paymentID, payerID string) error {
	_, err := s.client.ExecuteApprovedPayment(paymentID, payerID)
	if err != nil {
		s.logger.Println(err.Error())
		return err
	}
	return nil
}
