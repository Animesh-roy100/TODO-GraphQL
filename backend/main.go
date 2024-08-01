package main

import (
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/resolvers"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &resolvers.Schema,
		GraphiQL: true,
	})

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	// Wrap your GraphQL handler with the CORS handler
	http.Handle("/graphql", c.Handler(h))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
