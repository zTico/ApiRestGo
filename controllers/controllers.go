package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	for _, personalitie := range models.Personalities {
		if strconv.Itoa(personalitie.Id) == id {
			json.NewEncoder(w).Encode(personalitie)
		}
	}
}
