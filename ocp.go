package main

import (
	"errors"
)

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	CalculateAnnualInterest() float64
	GetBalance() float64
	GetAccountStatus() bool
}

type InvestmentAccount struct {
	balance float64
	locked  bool
}

func (account *InvestmentAccount) Deposit(amount float64) {
	account.balance += amount
}

func (account *InvestmentAccount) Withdraw(amount float64) error {
	locked := account.GetAccountStatus()
	if locked {
		return errors.New("Account is locked and cannot be withdrawed until certain period.")
	}

	if account.balance < amount {
		return errors.New("Insufficient balance.")
	}

	account.balance -= amount
	return nil
}

func (account *InvestmentAccount) CalculateAnnualInterest() float64 {
	return account.balance * (float64(5) / 10)
}

func (account *InvestmentAccount) GetBalance() float64 {
	return account.balance
}

func (account *InvestmentAccount) GetAccountStatus() bool {
	return account.locked
}

type SavingAccount struct {
}

func NewAccount(accountType string, locked bool) BankAccount {
	switch accountType {
	case "investment":
		return &InvestmentAccount{locked: locked}
	case "savings":
		return nil
	default:
		return nil
	}
}
