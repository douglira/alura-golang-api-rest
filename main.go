package main

import (
	"fmt"

	"github.com/douglira/alura-golang-api-rest.git/database"
	"github.com/douglira/alura-golang-api-rest.git/routes"
)

func main() {
	database.InitiateConnection()
	fmt.Println("Initializing Go Server")
	routes.HandleRequest()
}
