package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/felipealvesprestes/rest-api-roots/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Personalidades)
}
