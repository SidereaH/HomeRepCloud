package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageName   string `json:"image_name" gorm:"unique;not null"`
	PathToFile  string `json:"path_to_file" gorm:"not null"`
	Size        int64  `json:"size"`
	Category    string `json:"category"`
	Description string `json:"description"`
}
