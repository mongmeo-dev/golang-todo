package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mongmeo.dev/todo/internal/adapter/database/ent"
	"os"
)

func main() {
	client, err := ent.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("An exception occurred while connecting to mysql server %s", os.Getenv("DB_DSN"))
	}
	defer client.Close()
}
