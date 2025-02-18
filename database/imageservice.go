package database

import (
	"HomeRepCloud/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitImages(db *gorm.DB) {
	images := []models.Image{
		{
			ImageName:   "Электрика",
			PathToFile:  "./files/images/categories/electricity_category.png",
			Size:        0,
			Category:    "image",
			Description: "Икнока под электрику",
		},
		{
			ImageName:   "Газовое_оборудование",
			PathToFile:  "./files/images/categories/gas_category.png",
			Size:        0,
			Category:    "image",
			Description: "Икнока под газовое оборудование",
		},
		{
			ImageName:   "Бытовая_техника",
			PathToFile:  "./files/images/categories/household_appliances_category.png",
			Size:        0,
			Category:    "image",
			Description: "Икнока под бытовую технику",
		},
		{
			ImageName:   "Сантехника",
			PathToFile:  "./files/images/categories/plumbing_category.png",
			Size:        0,
			Category:    "image/category",
			Description: "Икнока под сантехнику",
		},
		//иконки
	}
	insertSizes(images)
	createImageInstances(db, images)

}
func createImageInstances(db *gorm.DB, images []models.Image) {
	for _, image := range images {
		// Используем FirstOrCreate для предотвращения дублирования записей
		result := db.FirstOrCreate(&image, models.Image{
			ImageName:   image.ImageName,
			PathToFile:  image.PathToFile,
			Size:        image.Size,
			Description: image.Description})
		if result.Error != nil {
			log.Println("Error while creating group", result.Error)
		}
	}
}
func insertSizes(images []models.Image) {
	for i := range images { // Используем индекс для изменения оригинального элемента
		images[i].Size = getImageSizeFromFilePath(images[i].PathToFile)
	}
}
func GetImageByID(id uint) (models.Image, error) {
	var image models.Image
	err := Instance.Where("id = ?", id).Find(&image).Error
	if err != nil {
		return models.Image{}, err
	}
	return image, nil
}
func GetImages() ([]models.Image, error) {
	var images []models.Image
	err := Instance.Find(&images).Error
	if err != nil {
		return []models.Image{}, err
	}
	return images, nil
}
func GetImageByName(name string) (models.Image, error) {
	var image models.Image
	err := Instance.Where("image_name = ?", name).First(&image).Error
	if err != nil {
		return models.Image{}, err
	}
	return image, nil
}
func getImageSizeFromImage(image models.Image) int64 {
	pathToFile := image.PathToFile
	fi, err := os.Stat(pathToFile)
	if err != nil {
		// Could not obtain stat, handle error
	}
	log.Printf("The file %s is %d bytes long", pathToFile, fi.Size())
	return fi.Size()
}
func getImageSizeFromFilePath(filePath string) int64 {
	fi, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	log.Printf("The file is %d bytes long", fi.Size())
	return fi.Size()
}
