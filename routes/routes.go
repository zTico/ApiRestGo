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
	route.HandleFunc("/api/personalidade", controllers.CreateNewData).Methods("Post")
	route.HandleFunc("/api/personalidade/{id}", controllers.DeletePersonalitie).Methods("Delete")
	route.HandleFunc("/api/personalidade/{id}", controllers.EditPersonalitie).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", route))
}
