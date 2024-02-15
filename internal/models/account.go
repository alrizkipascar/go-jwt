package models

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int       `json:"id"`
	Email             string    `json:"email"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	EncryptedPassword string    `json:"-"`
	Balance           int64     `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}

type CreateAccountRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

func (a *Account) NewAccount(email, firstName, lastName, password string) error {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Email = email
	a.FirstName = firstName
	a.LastName = lastName
	a.EncryptedPassword = string(encpw)
	a.Number = int64(rand.Intn(1000000))
	a.CreatedAt = time.Now().UTC()
	return nil
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}
