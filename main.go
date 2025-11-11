package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Ej0416/GoRSSAggregrator/internal/config"
	"github.com/Ej0416/GoRSSAggregrator/handlers"
	"github.com/Ej0416/GoRSSAggregrator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not fund in environment variables")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL not fund in environment variables")
	}

	dbConn, dbConErr := sql.Open("postgres",dbUrl)
	if dbConErr != nil {
		log.Fatal("unable to connect to database:", dbConErr)
	}

	queries := database.New(dbConn);
	apiCfg := config.ApiConfig{
		DB: queries,
	}

	fmt.Println(apiCfg)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlers.ReadinessHandler)
	v1Router.Get("/err", handlers.ErrorHandler)
	ext := &handlers.ExtendedApiConfig{ApiConfig: &apiCfg}
	v1Router.Post("/users", ext.CreateUserHandler)

	router.Mount("/v1",v1Router)

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
