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
	p := paypalsdk.Payment{
		Intent: "sale",
		Payer: &paypalsdk.Payer{
			PaymentMethod: "paypal",
		},
		Transactions: []paypalsdk.Transaction{paypalsdk.Transaction{
			Amount: &paypalsdk.Amount{
				Currency: request.Currency,
				Total:    request.Total,
			},
			Description: "My Payment",
		}},
		RedirectURLs: &paypalsdk.RedirectURLs{
			ReturnURL: "http://localhost:8080/v1/paypal/execute",
			CancelURL: "http://exemple.com/cancel",
		},
	}
	payment, err := s.client.CreatePayment(p)
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
		s.logger.Fatal(err)
		return err
	}
	return nil
}
