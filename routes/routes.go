package routes

import (
	"log"
	"net/http"

	"github.com/zTico/ApiRestGo/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personalidades", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
