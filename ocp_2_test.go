package main

import "testing"

/*
Payment Processing System with Single Responsibility Principle and Open-Closed Principle
Checked and reviewed by ChatGPT
*/

func TestCreditCardPaymentMethod_ProcessPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor(CreditCard)
	err := paymentProcessor.ProcessPayment(150000)
	if err != nil {
		t.Errorf("[TestCreditCardPaymentMethod_ProcessPayment] %v", err)
	}
}

func TestCreditCardPaymentMethod_RefundPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor(CreditCard)
	err := paymentProcessor.RefundPayment(150000)
	if err != nil {
		t.Errorf("[TestCreditCardPaymentMethod_RefundPayment] %v", err)
	}
}

func TestPayPalPaymentMethod_ProcessPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor(PayPal)
	err := paymentProcessor.ProcessPayment(150000)
	if err != nil {
		t.Errorf("[TestPayPalPaymentMethod_ProcessPayment] %v", err)
	}
}

func TestPayPalPaymentMethod_RefundPayment(t *testing.T) {
	paymentProcessor := NewPaymentProcessor(PayPal)
	err := paymentProcessor.RefundPayment(150000)
	if err != nil {
		t.Errorf("[TestPayPalPaymentMethod_RefundPayment] %v", err)
	}
}
