package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("oi")
	HandleRequest()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
