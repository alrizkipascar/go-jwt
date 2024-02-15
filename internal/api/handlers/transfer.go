package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/models"
	"github.com/alrizkipascar/go-jwt/internal/utils"
)

func Transfer(w http.ResponseWriter, r *http.Request) error {
	TransferReq := new(models.TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(TransferReq); err != nil {
		return err
	}
	defer r.Body.Close()
	return utils.WriteJSON(w, http.StatusOK, TransferReq)
}
