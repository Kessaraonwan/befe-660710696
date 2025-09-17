package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

var db *sql.DB

func initDB() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "postgres")
	name := getEnv("DB_NAME", "bookstore")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database (check host/port/user/pass/dbname): %v", err)
	}

	fmt.Println("✅ Connected to Postgres OK")
}

func main() {
	// ต่อ DB ก่อน
	initDB()
	defer db.Close()

	r := gin.Default()

	// health check
	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "unhealthy",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	// รันที่พอร์ต 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
