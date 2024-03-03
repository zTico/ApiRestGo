package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zTico/ApiRestGo/database"
	"github.com/zTico/ApiRestGo/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personalitie

	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
}

func ReturnOnePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalitie models.Personalitie

	database.DB.First(&personalitie, id)

	json.NewEncoder(w).Encode(personalitie)
}

func CreateNewData(w http.ResponseWriter, r *http.Request) {
	var newPersonalidade models.Personalitie

	json.NewDecoder(r.Body).Decode(&newPersonalidade)

	database.DB.Create(&newPersonalidade)

	json.NewEncoder(w).Encode(newPersonalidade)
}

func DeletePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalitie models.Personalitie

	database.DB.Delete(&personalitie, id)

	json.NewEncoder(w).Encode(personalitie)
}
