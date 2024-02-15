package handlers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/auth"
	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/models"
	"github.com/alrizkipascar/go-jwt/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	// if r.Method != "POST" {
	// 	return fmt.Errorf("method not allowed %s", r.Method)
	// }
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	acc, err := database.GetAccountByEmail(req.Email)

	// acc, err := database.GetAccountByNumber(int(req.Number))
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}
	token, err := auth.CreateJWT(acc)
	if err != nil {
		return err
	}

	resp := models.LoginResponseEmail{
		Token: token,
		Email: req.Email,
	}

	fmt.Printf("%+v\n", acc)
	return utils.WriteJSON(w, http.StatusOK, resp)
}

// EMAIL COMMENT

func LoginWithEmail(w http.ResponseWriter, r *http.Request) error {
	// if r.Method != "POST" {
	// 	return fmt.Errorf("method not allowed %s", r.Method)
	// }
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := database.GetAccountByEmail(req.Email)
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}
	token, err := auth.CreateJWT(acc)
	if err != nil {
		return err
	}

	resp := models.LoginResponseEmail{
		Token: token,
		Email: req.Email,
	}

	fmt.Printf("%+v\n", acc)
	return utils.WriteJSON(w, http.StatusOK, resp)
}

// func CreateJWT(account *models.Account) (string, error) {
// 	// Create the Claims
// 	claims := &jwt.MapClaims{
// 		"expiresAt":     15000,
// 		"accountNumber": account.Number,
// 	}

// 	secret := os.Getenv("JWT_SECRET")
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// ss, err := token.SignedString(mySigningKey)
// 	// fmt.Printf("%v %v", ss, err)
// 	return token.SignedString([]byte(secret))
// }

// func permissionDenied(w http.ResponseWriter) {
// 	utils.WriteJSON(w, http.StatusForbidden, middlewares.ApiError{Error: "permission denied"})
// }

// // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50TnVtYmVyIjo3MDA0NTkzLCJleHBpcmVzQXQiOjE1MDAwfQ.TuIPQuYpGAeRKJ92akkYI4oTPR-ymrFJ-fNWMrno9M0

// func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("calling JWT auth mmiddleware")
// 		tokenString := r.Header.Get("x-jwt-token")
// 		token, err := validateJWT(tokenString)
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}

// 		if !token.Valid {
// 			permissionDenied(w)
// 			return
// 		}

// 		userID, err := GetID(r)
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}

// 		account, err := database.GetAccountByID(userID)
// 		// account, err := s.GetUserByJWTToken(userID)
// 		// // err
// 		if err != nil {
// 			permissionDenied(w)
// 			return
// 		}

// 		claims := token.Claims.(jwt.MapClaims)
// 		// panic(reflect.TypeOf(claims["accountNumber"]))
// 		if account.Number != int64(claims["accountNumber"].(float64)) {
// 			permissionDenied(w)
// 			return
// 		}

// 		if err != nil {
// 			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
// 			return
// 		}
// 		// // account/{id}

// 		// // account, err :=

// 		// claims := token.Claims.(jwt.MapClaims)
// 		// if claims["ID"] == account.ID
// 		// fmt.Println(claims)

// 		handlerFunc(w, r)
// 	}
// }

// func validateJWT(tokenString string) (*jwt.Token, error) {
// 	secret := os.Getenv("JWT_SECRET")
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(secret), nil
// 	})
// }
