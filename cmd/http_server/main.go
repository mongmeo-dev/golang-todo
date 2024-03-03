package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mongmeo.dev/todo/internal/adapter/database/ent"
	"net/http"
	"os"
)

func main() {
	client, err := ent.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("An exception occurred while connecting to mysql server %s", os.Getenv("DB_DSN"))
	}
	defer client.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Fail to start application")
	}
}
