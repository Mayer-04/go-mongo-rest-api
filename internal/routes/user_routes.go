package routes

import (
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/controllers"
)

func Routes(router *http.ServeMux) *http.ServeMux {

	router.HandleFunc("GET /users", controllers.GetUsers)

	router.HandleFunc("GET /users/{id}", controllers.GetUserByID)

	router.HandleFunc("POST /users", controllers.CreateUser)

	router.HandleFunc("PUT /users/{id}", controllers.UpdateUser)

	router.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)

	return router

}
