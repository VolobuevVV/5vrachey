package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-app/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("Connected to database successfully!")
}

func runMigrations() {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating migration driver:", err)
	}

	migrationsPath := "file://migrations"
	if _, err := os.Stat("/app/backend/migrations"); err == nil {
		migrationsPath = "file:///app/backend/migrations" // в контейнере
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatal("Error creating migration instance:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}

	log.Println("Migrations applied successfully!")
}

func main() {
	files, _ := os.ReadDir("/app/backend/migrations")
	log.Printf("Files in /app/backend/migrations:")
	for _, file := range files {
		log.Printf(" - %s", file.Name())
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	InitDB()
	runMigrations()

	router := gin.Default()

	promotionHandler := &handlers.PromotionHandler{DB: DB}
	router.GET("/api/top-panel-text", promotionHandler.GetPromotionText)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
