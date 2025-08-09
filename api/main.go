package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LuizFernando991/golang-api/api/router"
	"github.com/LuizFernando991/golang-api/infra/config"
)

func main() {
	envVariables, err := config.LoadEnv(".")
	if err != nil {
		log.Fatalf("Variables error: %v", err)
	}

	router := router.InitializeRouter()

	port := ":" + envVariables.API_PORT
	fmt.Printf("Server running on %s\n", envVariables.API_PORT)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Erro crítico: %v", err)
	}
}
