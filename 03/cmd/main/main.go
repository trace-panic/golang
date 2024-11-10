package main

import (
	"03/pkg/config"
	"03/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)

	http.Handle("/", r)
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
