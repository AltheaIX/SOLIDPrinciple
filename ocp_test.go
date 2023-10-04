package main

import (
	"testing"
)

func TestInvestmentAccount_Deposit(t *testing.T) {
	investment := NewAccount("investment", true)
	investment.Deposit(5000)
	t.Logf("[TestInvestmentAccount_Deposit] %v", investment.GetBalance())
}

func TestSavingsAccount_Deposit(t *testing.T) {
	savings := NewAccount("savings", false)
	savings.Deposit(150000)
	t.Logf("[TestSavingsAccount_Deposit] %v", savings.GetBalance())
}

func TestInvestmentAccount_Withdraw(t *testing.T) {
	investment := NewAccount("investment", true)
	err := investment.Withdraw(5000)
	if err != nil {
		t.Errorf("[TestInvestmentAccount_Withdraw] %v - ERROR", err)
	}

	t.Logf("[TestInvestmentAccount_Withdraw] %v", investment.GetBalance())
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	savings := NewAccount("savings", true)
	savings.Deposit(502123)
	err := savings.Withdraw(5000)
	if err != nil {
		t.Errorf("[TestSavingsAccount_Withdraw] %v - ERROR", err)
	}

	t.Logf("[TestSavingsAccount_Withdraw] %v", savings.GetBalance())
}

func TestInvestmentAccount_CalculateAnnualInterest(t *testing.T) {
	investment := NewAccount("investment", true)
	investment.Deposit(50000)

	t.Logf("[TestInvestmentAccount_CalculateAnnualInterest] You will be getting Rp %v annually.", investment.CalculateAnnualInterest())
}

func TestSavingsAccount_CalculateAnnualInterest(t *testing.T) {
	savings := NewAccount("savings", true)
	savings.Deposit(50000)

	t.Logf("[TestSavingsAccount_CalculateAnnualInterest] You will be getting Rp %v annually.", savings.CalculateAnnualInterest())
}

func TestInvestmentAccount_GetAccountStatus(t *testing.T) {
	investment := NewAccount("investment", true)
	t.Logf("[TestInvestmentAccount_GetAccountStatus] %v", investment.GetAccountStatus())
}

func TestSavingsAccount_GetAccountStatus(t *testing.T) {
	savings := NewAccount("savings", true)
	t.Logf("[TestSavingsAccount_GetAccountStatus] %v", savings.GetAccountStatus())
}
