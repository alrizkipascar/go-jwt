package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID                int       `json:"id"`
	UUID              uuid.UUID `json:"userId"`
	Email             string    `json:"email"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	Activated         int64     `json:"activated"`
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

// Testing for certain user.
type UserAccount struct {
	// ID        int       `json:"id"`
	UUID      uuid.UUID `json:"userId"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	// Activated int64     `json:"activated"`
	Balance int64 `json:"balance"`
}

func (a *Account) NewAccount(email, firstName, lastName, password string) error {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Email = email
	a.UUID = uuid.New()
	a.FirstName = firstName
	a.LastName = lastName
	a.Activated = 0
	a.EncryptedPassword = string(encpw)
	a.Number = int64(rand.Intn(1000000))
	a.CreatedAt = time.Now().UTC()
	return nil
}

func (a *Account) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(pw)) == nil
}
