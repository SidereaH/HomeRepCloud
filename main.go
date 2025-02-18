package main

import (
	"HomeRepCloud/controllers"
	"HomeRepCloud/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	database.Connect("host=localhost user=postgres password=postgres dbname=home_rep_cloud port=5433 sslmode=disable TimeZone=Europe/Moscow")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run("0.0.0.0:8081")
}
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
	return router
}
