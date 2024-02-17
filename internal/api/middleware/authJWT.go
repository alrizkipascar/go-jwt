package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/auth"
	"github.com/alrizkipascar/go-jwt/internal/database"
	"github.com/alrizkipascar/go-jwt/internal/helpers"
	"github.com/alrizkipascar/go-jwt/internal/utils"
	"github.com/golang-jwt/jwt/v4"
)

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")

		// Ensure that user and x-jwt-token is valid
		tokenString := r.Header.Get("x-jwt-token")
		userId := r.Header.Get("user-id")

		token, err := auth.ValidateJWT(tokenString)
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}
		if !token.Valid {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		userID, err := helpers.ConvertHeaderID(userId)
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		account, err := database.GetAccountByID(userID)

		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		// previous -> accountNumber
		if account.Email != claims["accountEmail"] {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}
		fmt.Println("3", account)

		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}

		handlerFunc(w, r)
	}
}
