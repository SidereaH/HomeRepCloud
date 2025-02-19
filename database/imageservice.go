package database

import (
	"HomeRepCloud/models"
	"encoding/json"
	"log"
	"os"
)

func InitImages() {
	images := []models.Image{
		{
			ImageName:   "Электрика",
			PathToFile:  "./files/images/categories/electricity_category.png",
			Size:        getImageSizeFromFilePath("./files/images/categories/electricity_category.png"),
			Category:    "image",
			Description: "Иконка под электрику",
		},
		{
			ImageName:   "Газовое_оборудование",
			PathToFile:  "./files/images/categories/gas_category.png",
			Size:        getImageSizeFromFilePath("./files/images/categories/gas_category.png"),
			Category:    "image",
			Description: "Иконка под газовое оборудование",
		},
		{
			ImageName:   "Бытовая_техника",
			PathToFile:  "./files/images/categories/household_appliances_category.png",
			Size:        getImageSizeFromFilePath("./files/images/categories/household_appliances_category.png"),
			Category:    "image",
			Description: "Икнока под бытовую технику",
		},
		{
			ImageName:   "Сантехника",
			PathToFile:  "./files/images/categories/plumbing_category.png",
			Size:        getImageSizeFromFilePath("./files/images/categories/plumbing_category.png"),
			Category:    "image/category",
			Description: "Икнока под сантехнику",
		},
	}

	for _, image := range images {
		saveImageToRedis(image)
	}
	log.Println("Images initialized in Redis!")
}

func saveImageToRedis(image models.Image) {
	data, err := json.Marshal(image)
	if err != nil {
		log.Println("Error marshalling image: ", err)
		return
	}
	RedisClient.Set(ctx, image.ImageName, data, 0)
}

func GetImageByName(name string) (models.Image, error) {
	data, err := RedisClient.Get(ctx, name).Result()
	if err != nil {
		return models.Image{}, err
	}
	var image models.Image
	json.Unmarshal([]byte(data), &image)
	return image, nil
}

func GetImages() ([]models.Image, error) {
	keys, err := RedisClient.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	var images []models.Image
	for _, key := range keys {
		data, err := RedisClient.Get(ctx, key).Result()
		if err == nil {
			var image models.Image
			json.Unmarshal([]byte(data), &image)
			images = append(images, image)
		}
	}
	return images, nil
}

func getImageSizeFromFilePath(filePath string) int64 {
	fi, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fi.Size()
}
