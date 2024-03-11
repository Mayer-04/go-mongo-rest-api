package routes

import (
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/controllers"
)

func Routes(app *http.ServeMux) *http.ServeMux {

	app.HandleFunc("GET /", controllers.GetUsers)

	app.HandleFunc("GET /{id}", controllers.GetUser)

	app.HandleFunc("POST /", controllers.CreateUser)

	app.HandleFunc("PUT /{id}", controllers.UpdateUser)

	app.HandleFunc("DELETE /{id}", controllers.DeleteUser)

	return app

}
