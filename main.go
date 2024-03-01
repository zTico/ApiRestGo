package main

import (
	"fmt"

	"github.com/zTico/ApiRestGo/models"
	"github.com/zTico/ApiRestGo/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Id: 1, Name: "Nome 1", History: "Historia 1"},
		{Id: 2, Name: "Nome 2", History: "Historia 2"},
	}

	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
