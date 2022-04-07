package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"rest-api-roots/database"
	"rest-api-roots/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func RetornoPersonalidadeEspecifica(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
