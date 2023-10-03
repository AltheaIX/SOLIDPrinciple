package main

import (
	"testing"
)

func TestInvestmentAccount_Deposit(t *testing.T) {
	investment := NewAccount("investment", true)
	investment.Deposit(5000)
	t.Logf("[TestInvestmentAccount_Deposit] %v", investment.GetBalance())
}

func TestInvestmentAccount_Withdraw(t *testing.T) {
	/*	investment := NewAccount("investment", true)
		err := investment.Withdraw(5000)
		if err != nil {
			t.Logf("[TestInvestmentAccount_Withdraw] %v - ERROR", err)
		}*/

	investment := NewAccount("savings", false)
	err := investment.Withdraw(5000)
	if err != nil {
		t.Logf("[TestInvestmentAccount_Withdraw] %v - ERROR", err)
	}
	t.Logf("[TestInvestmentAccount_Withdraw] %v", investment.GetBalance())
}

func TestInvestmentAccount_CalculateAnnualInterest(t *testing.T) {
	investment := NewAccount("investment", true)
	investment.Deposit(50000)
	t.Log(investment.CalculateAnnualInterest())

	t.Logf("[TestInvestmentAccount_CalculateAnnualInterest] You will be getting Rp %v annually.", investment.CalculateAnnualInterest())
}

func TestInvestmentAccount_GetAccountStatus(t *testing.T) {
	investment := NewAccount("investment", true)
	t.Logf("[TestInvestmentAccount_GetAccountStatus] %v", investment.GetAccountStatus())
}
