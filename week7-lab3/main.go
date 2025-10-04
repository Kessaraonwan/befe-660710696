package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB() {
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conST := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, name)
	//t.Println(conSt)
	db, err = sql.Open("postgres", conST)
	if err != nil {
		log.Fatal("failed to open detabase")
	}
}

func main() {
	initDB()
}
