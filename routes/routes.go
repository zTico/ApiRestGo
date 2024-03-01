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
	route.HandleFunc("/personalidades", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", route))
}
