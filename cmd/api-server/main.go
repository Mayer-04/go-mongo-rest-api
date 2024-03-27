package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Mayer-04/go-mongo-rest-api/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	app := http.NewServeMux()

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	userRoutes := routes.Routes(app)

	app.Handle("/api/", http.StripPrefix("/api", userRoutes))

	log.Println("Server is running on port:", port)

	err := http.ListenAndServe("localhost:"+port, app)

	if err != nil {
		panic(err)
	}
}
