package main

import (
	"errors"
)

/*
Bank Account Management Support Multiple Account And Locked System with Single Responsibility Principle and Open-Closed Principle
Checked and reviewed by ChatGPT
*/

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	CalculateAnnualInterest() float64
	GetBalance() float64
	GetAccountStatus() bool
}

type InterestCalculator interface {
	CalculateAnnualInterest(balance float64) float64
}

type InvestmentInterestCalculator struct{}

func (calc *InvestmentInterestCalculator) CalculateAnnualInterest(balance float64) float64 {
	return balance * (float64(5) / 100)
}

type InvestmentAccount struct {
	balance            float64
	locked             bool
	interestCalculator InterestCalculator
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
	return account.interestCalculator.CalculateAnnualInterest(account.balance)
}

func (account *InvestmentAccount) GetBalance() float64 {
	return account.balance
}

func (account *InvestmentAccount) GetAccountStatus() bool {
	return account.locked
}

type SavingsInterestCalculator struct{}

func (calc *SavingsInterestCalculator) CalculateAnnualInterest(balance float64) float64 {
	return balance * (float64(0.5) / 100)
}

type SavingsAccount struct {
	balance            float64
	interestCalculator InterestCalculator
}

func (account *SavingsAccount) Deposit(amount float64) {
	account.balance += amount
}

func (account *SavingsAccount) Withdraw(amount float64) error {
	if account.balance < amount {
		return errors.New("Insufficient balance.")
	}

	account.balance -= amount
	return nil
}

func (account *SavingsAccount) CalculateAnnualInterest() float64 {
	return account.interestCalculator.CalculateAnnualInterest(account.balance)
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}

func (account *SavingsAccount) GetAccountStatus() bool {
	return false
}

func NewAccount(accountType string, locked bool) BankAccount {
	switch accountType {
	case "investment":
		return &InvestmentAccount{locked: locked, interestCalculator: &InvestmentInterestCalculator{}}
	case "savings":
		return &SavingsAccount{interestCalculator: &SavingsInterestCalculator{}}
	default:
		return nil
	}
}
