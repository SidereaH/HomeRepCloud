package controllers

import (
	"HomeRepCloud/database"
	"HomeRepCloud/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"

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
	category := context.PostForm("category")
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка загрузки файла"})
		return
	}
	defer file.Close()

	dir := filepath.Join("files", "images", category)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания директории"})
		return
	}

	filePath := filepath.Join(dir, header.Filename)

	// Проверка, существует ли уже изображение в Redis
	exists, err := database.RedisClient.Exists(context, header.Filename).Result()
	if err == nil && exists > 0 {
		context.JSON(http.StatusConflict, gin.H{"error": "Файл уже существует"})
		return
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения файла"})
		return
	}
	defer outFile.Close()

	size := header.Size
	imageData := models.Image{
		ImageName:   header.Filename,
		PathToFile:  filePath,
		Size:        size,
		Category:    category,
		Description: context.PostForm("description"),
	}

	// Сохранение информации в Redis
	database.RedisClient.HSet(context, header.Filename, map[string]interface{}{
		"image_name":   imageData.ImageName,
		"path_to_file": imageData.PathToFile,
		"size":         imageData.Size,
		"category":     imageData.Category,
		"description":  imageData.Description,
	})

	context.JSON(http.StatusOK, gin.H{"message": "Файл успешно загружен", "path": filePath})
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
