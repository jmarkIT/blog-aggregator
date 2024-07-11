package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/blog-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DBURL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	dbQueries := database.New(db)

	apiConfig := apiConfig{DB: dbQueries}

	m := http.NewServeMux()
	m.HandleFunc("/v1/healthz", handlerReadiness)
	m.HandleFunc("/v1/err", handlerError)
	m.HandleFunc("POST /v1/users", apiConfig.createUserHandler)
	s := &http.Server{
		Addr:    ":" + port,
		Handler: m,
	}
	log.Fatal(s.ListenAndServe())
}
