package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/models"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func seedAccount(email, fname, lname, pw string) *models.Account {
	acc := new(models.Account)
	err := acc.NewAccount(email, fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func SeedAccounts() {
	seedAccount("alrizkipasca@gmail.com", "alrizki", "pasca", "pasca99")
}
