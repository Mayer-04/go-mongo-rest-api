package routes

import (
	"net/http"

	"github.com/Mayer-04/go-mongo-rest-api/internal/controllers"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET api/user/", controllers.GetUsers)

	router.HandleFunc("GET api/user/{id}", controllers.GetUser)

	router.HandleFunc("POST api/user/", controllers.CreateUser)

	router.HandleFunc("PUT api/user/{id}", controllers.UpdateUser)

	router.HandleFunc("DELETE api/user/{id}", controllers.DeleteUser)

	return router
}
