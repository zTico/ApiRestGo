package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zTico/ApiRestGo/controllers"
)

func HandleRequest() {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Home)
	route.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	route.HandleFunc("/api/personalidades/{id}", controllers.ReturnOnePersonalitie).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", route))
}
