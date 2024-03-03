package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mongmeo.dev/todo/internal/adapter/database/ent"
	TodoRouter "mongmeo.dev/todo/internal/adapter/router/todo"
	TodoApplication "mongmeo.dev/todo/internal/application/todo"
	"net/http"
	"os"
)

func main() {
	client, err := ent.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("An exception occurred while connecting to mysql server %s", os.Getenv("DB_DSN"))
	}
	defer client.Close()

	todoApplication := TodoApplication.New(client.Todo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/", TodoRouter.New(todoApplication).GetRouter())

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Fail to start application")
	}
}
