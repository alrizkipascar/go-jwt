package helpers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get ID by URL
func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

// Get ID by HEADER
func ConvertHeaderID(r string) (int, error) {
	id, err := strconv.Atoi(r)
	if err != nil {
		return id, fmt.Errorf("invalid id given %v", id)
	}
	return id, nil
}
