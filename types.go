package main

import "math/rand"

type Account struct {
ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string`json:"last_name"`
	BankID int64`json:"bankID"`
	Balance int64`json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(1000),
		FirstName: firstName,
		LastName: lastName,
		BankID: int64(rand.Intn(100000000)),
	}
}
