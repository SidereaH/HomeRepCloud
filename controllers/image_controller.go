package controllers

import (
	"HomeRepCloud/database"
	"HomeRepCloud/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type ImageController struct {
	Type  string
	Image models.Image
}

func SaveImage(context *gin.Context) {

}
func GetAvailableGroups(context *gin.Context) {
	//запрос к спрингу которого пока не существуе
}
func GetImageByCategory(context *gin.Context) {

}

// по описанию типа сантхеника
func GetImageByName(c *gin.Context) {
	imageName := c.Param("name")
	fmt.Println(imageName)
	image, err := database.GetImageByName(imageName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	data, err := os.ReadFile(image.PathToFile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		c.Abort()
		return
	}
	c.Data(http.StatusOK, "image/png", data)

}
func GetAllImages(context *gin.Context) {
	images, err := database.GetImages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"images": images})
}
