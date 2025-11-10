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
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not fund in environment variables")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Server is starting on port:", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on port:", port)
}
