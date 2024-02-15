package handlers

import (
	// "crypto/rand"
	"encoding/json"
	"fmt"

	// "log"
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/auth"
	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/helpers"
	"github.com/alrizkipascar/go-jwt/internal/models"
	"github.com/alrizkipascar/go-jwt/internal/utils"
	// "os"
	// "strconv"
	// jwt "github.com/golang-jwt/jwt/v4"
	// "github.com/gorilla/mux"
)

func GetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := database.GetAccounts()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, accounts)
}

func GetAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := helpers.GetID(r)
		if err != nil {
			return err
		}
		account, err := database.GetAccountByID(id)
		if err != nil {
			return err
		}
		return utils.WriteJSON(w, http.StatusOK, account)
	}
	if r.Method == "DELETE" {
		return DeleteAccount(w, r)
	}
	return fmt.Errorf("Method not allowed %s", r.Method)
	// account := NewAccount("Anthony", "GG")
	// fmt.Println(id)
	// return WriteJSON(w, http.StatusOK, &Account{})
}

func CreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	account := new(models.Account)
	err := account.NewAccount(req.Email, req.FirstName, req.LastName, req.Password)

	if err != nil {
		return err
	}

	if err := database.CreateAccount(account); err != nil {
		return err
	}

	tokenString, err := auth.CreateJWT(account)
	if err != nil {
		return err
	}
	fmt.Println("JWT token:", tokenString)
	return utils.WriteJSON(w, http.StatusOK, account)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}
	if err := database.DeleteAccount(id); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

// func NewAccount(email, firstName, lastName, password string) (*models.Account, error) {
// 	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Account{
// 		Email:             email,
// 		FirstName:         firstName,
// 		LastName:          lastName,
// 		EncryptedPassword: string(encpw),
// 		Number:            int64(rand.Intn(1000000)),
// 		CreatedAt:         time.Now().UTC(),
// 	}, nil
// }
