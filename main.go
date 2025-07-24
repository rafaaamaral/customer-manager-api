package main

import (
	"customer-manager-api/src/config"
	"customer-manager-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Println("Starting Customer Manager API...")

	r := router.SetupRouter()

	fmt.Printf("Customer Manager API is running on port %d\n", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
