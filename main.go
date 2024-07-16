package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	welcome := "Welcome to management "

	route := mux.NewRouter()
	route.HandleFunc("/test", TestServer).Methods("GET")
	server := &http.Server{
		Handler:      route,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf(welcome)

	log.Fatal(server.ListenAndServe())

}

func TestServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))


}
