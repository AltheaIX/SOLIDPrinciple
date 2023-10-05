package main

import "testing"

func TestCreditCardPaymentMethod_ProcessPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor("creditcard")
	err := paymentProcessor.ProcessPayment(150000)
	if err != nil {
		t.Errorf("[TestCreditCardPaymentMethod_ProcessPayment] %v", err)
	}
}

func TestCreditCardPaymentMethod_RefundPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor("creditcard")
	err := paymentProcessor.RefundPayment(150000)
	if err != nil {
		t.Errorf("[TestCreditCardPaymentMethod_RefundPayment] %v", err)
	}
}

func TestPayPalPaymentMethod_ProcessPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor("paypal")
	err := paymentProcessor.ProcessPayment(150000)
	if err != nil {
		t.Errorf("[TestPayPalPaymentMethod_ProcessPayment] %v", err)
	}
}

func TestPayPalPaymentMethod_RefundPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor("paypal")
	err := paymentProcessor.RefundPayment(150000)
	if err != nil {
		t.Errorf("[TestPayPalPaymentMethod_RefundPayment] %v", err)
	}
}
