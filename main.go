package main

import (
	"fmt"

	"github.com/zTico/ApiRestGo/database"
	"github.com/zTico/ApiRestGo/routes"
)

func main() {
	database.ConectDataBase()

	fmt.Println("Iniciando servidor...")
	routes.HandleRequest()
}
