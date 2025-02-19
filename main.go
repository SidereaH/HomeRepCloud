package main

import (
	"HomeRepCloud/controllers"
	"HomeRepCloud/database"
	_ "HomeRepCloud/docs" // Импортируйте сгенерированную документацию
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func main() {
	database.Connect("localhost:6379", "", 0)
	database.InitImages()

	// Initialize Router
	router := initRouter()
	router.Run(":8081")
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
	}

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
