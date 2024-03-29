package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in the environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Database URL not found in the environment")
	}

	conn, err := sql.Open("postgres",dbUrl)

	if err != nil {
		log.Fatal("Databese connection failed: ", err)
	}
	db := database.New(conn)

	ApiConfig := apiConfig{
		DB: db,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()

	v1Router.Get("/status", handlerStatus)

	// User routes
	v1Router.Post("/user/register",ApiConfig.handlerRegisterUser)
	v1Router.Post("/user/login",ApiConfig.handlerLogin)

	// Stocks Routes
	v1Router.Post("/stock",ApiConfig.middlewareAuth(ApiConfig.handlerCreateStocks))
	v1Router.Get("/stock",ApiConfig.handlerGetStocks)
	v1Router.Put("/stock/{stockId}",ApiConfig.middlewareAuth(ApiConfig.handlerUpdateStocks))
	v1Router.Post("/stock/buy",ApiConfig.middlewareAuth(ApiConfig.handlerBuyStock))
	

	router.Mount("/api/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v\n", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
