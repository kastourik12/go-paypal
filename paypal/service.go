package paypal

import "fmt"

type Service struct {
	Client *CustomClient
}

func New(client CustomClient) *Service {
	return &Service{
		Client: &client,
	}
}

func (s *Service) CreatePayment(paymentRequest PaymentRequest) error {
	err := s.Client.CreatePayment(paymentRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
