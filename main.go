package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not specified")
	}

	fmt.Println("Port: ", portString)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		ExposedHeaders:   []string{"LINK"},
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1router)
	v1router.Get("/error", handlerError)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
