package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/jupitters/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){

	godotenv.Load(".env")
	
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT não encontrada no ambiente.")
	}

	dbURL := os.Getenv(("DB_URL"))
	if dbURL == "" {
		log.Fatal("DB_URL não encontrada no ambiente.")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options {
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	router.Mount("/v1", v1Router)
	v1Router.Get("/err", handlerErr)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+ portString,
	}

	log.Printf("Servidor iniciado em porta %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)
}