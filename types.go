package main

import "math/rand"

// defining how the variable will be serialized as when it gets serialized to JSON
type Account struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Number    int64   `json:"number"`
	Balance   float32 `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		ID:        rand.Intn(10000),
		Number:    int64(rand.Intn(1000000)),
		Balance:   0,
	}
}
