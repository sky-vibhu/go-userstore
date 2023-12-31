package routes

import (
	"github.com/gorilla/mux"
	"github.com/sky-vibhu/go-userstore/pkg/controllers"
)

var RegisterUserStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	//router.HandleFunc("/user/{userId}", controllers.GetUserById).Methods("GET")
	//router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
}