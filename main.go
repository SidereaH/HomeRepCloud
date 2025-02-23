package main

import (
	"HomeRepCloud/controllers"
	"HomeRepCloud/database"
	_ "HomeRepCloud/docs" // Импортируйте сгенерированную документацию
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := strconv.Atoi(os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal("Change DB name to int")
	}
	database.Connect(
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PASSWORD"),
		db,
	)
	database.InitImages()

	// Initialize Router
	router := initRouter()
	router.Run(os.Getenv("API_ADDRESS") + ":" + os.Getenv("API_PORT"))
}

// @title HomeRepCloud API
// @version 1.0
// @description API для работы с изображениями в HomeRepCloud.
// @host localhost:8081
// @BasePath /api

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			imagePath := "./files/images/hello.jpg"
			data, err := os.ReadFile(imagePath)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
				return
			}
			c.Data(http.StatusOK, "image/png", data)
		})

		api.GET("/images/category/:category", controllers.GetImageByCategory)
		api.GET("/image/:name", controllers.GetImageByName)
		api.GET("/images", controllers.GetAllImages)
		api.POST("/images/save", controllers.SaveImage)
	}

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
