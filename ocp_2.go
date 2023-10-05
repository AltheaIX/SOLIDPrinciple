package main

import "fmt"

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	RefundPayment(amount float64) error
}

type CreditCardPaymentMethod struct{}

func (c *CreditCardPaymentMethod) ProcessPayment(amount float64) error {
	fmt.Printf("[CreditCardProccessPayment] Your payment with due %v is on process...\n", amount)
	return nil
}

func (c *CreditCardPaymentMethod) RefundPayment(amount float64) error {
	fmt.Printf("[CreditCardRefundPayment] Your success payment with amount of %v has been refunded.\n", amount)
	return nil
}

type PayPalPaymentMethod struct{}

func (p *PayPalPaymentMethod) ProcessPayment(amount float64) error {
	fmt.Printf("[PayPalProcessPayment] Your payment with due %v is on process...\n", amount)
	return nil
}

func (p *PayPalPaymentMethod) RefundPayment(amount float64) error {
	fmt.Printf("[PayPalRefundPayment] Your success payment with amount of %v has been refunded.\n", amount)
	return nil
}

type PaymentProcessor struct {
	paymentMethod PaymentMethod
}

func (p *PaymentProcessor) ProcessPayment(amount float64) error {
	err := p.paymentMethod.ProcessPayment(amount)
	return err
}

func (p *PaymentProcessor) RefundPayment(amount float64) error {
	err := p.paymentMethod.RefundPayment(amount)
	return err
}

func NewPaymentProcessor(method string) PaymentMethod {
	switch method {
	case "creditcard":
		return &PaymentProcessor{paymentMethod: &CreditCardPaymentMethod{}}
	case "paypal":
		return &PaymentProcessor{paymentMethod: &PayPalPaymentMethod{}}
	default:
		return nil
	}
}
