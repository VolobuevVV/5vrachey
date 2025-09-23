package main

import (
	"log"
	"my-app/database"
	"my-app/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	database.InitDB()
	database.RunMigrations()

	router := gin.Default()

	promotionHandler := &handlers.PromotionHandler{DB: database.DB}
	router.GET("/api/top-panel-text", promotionHandler.GetPromotionText)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
