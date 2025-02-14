package controllers

import (
	"HomeRepCloud/models"
	"github.com/gin-gonic/gin"
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
func GetImageByGroup(context *gin.Context) {

}
