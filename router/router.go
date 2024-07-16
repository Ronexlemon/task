// router/routes.go
package router

import (
	"managementapi/controller"
	

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
    router := mux.NewRouter()

    // Define your routes here
    router.HandleFunc("/tasks",controller.Create ).Methods("POST")
	router.HandleFunc("/tasks",controller.Tasks).Methods("GET")
	router.HandleFunc("/tasks/{id}",controller.Delete).Methods("DELETE")
	router.HandleFunc("/tasks/{id}",controller.Update).Methods("PUT")
  

    return router
}


