package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Mayer-04/go-mongo-rest-api/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	routes.Routes()

	http.ListenAndServe("localhost:"+port, mux)
}
