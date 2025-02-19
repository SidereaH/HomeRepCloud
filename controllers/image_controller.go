package controllers

import (
	"HomeRepCloud/database"
	"HomeRepCloud/models"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/swaggo/swag/example/celler/httputil"
	//"github.com/swaggo/swag/example/celler/model"
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

// @Summary Получить изображение по названию
// @Description Возвращает изображение по указанному названию
// @Tags images
// @Accept json
// @Produce image/png
// @Param name path string true "Название изображения"
// @Success 200 {file} png "Изображение"
// @Failure 404 {object} map[string]string "Image not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /image/{name} [get]
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

// @Summary Получить все изображения
// @Description Возвращает список всех изображений
// @Tags images
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Список изображений"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /images [get]
func GetAllImages(context *gin.Context) {
	images, err := database.GetImages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"images": images})
}
