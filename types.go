package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	toAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Task struct {
	ID         int       `json:"id"`
	TaskName   string    `json:"firstName"`
	Status     int       `json:"status"`
	TimePosted time.Time `json:"timeposted"`
	// SubTask  []SubTask
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		// ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(10000000)),
		CreatedAt: time.Now().UTC(),
	}
}
