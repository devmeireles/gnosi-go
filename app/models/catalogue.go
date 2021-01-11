package models

import "github.com/jinzhu/gorm"

// Catalogue struct
type Catalogue struct {
	gorm.Model
	Title       string  `json:"title" gorm:"unique;not null"`
	Slug        string  `json:"slug" gorm:"unique;not null"`
	Description string  `json:"description" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	MediaID     int     `json:"media_id"`
	Status      int     `json:"status" sql:"DEFAULT:1"`
	// Seasons     []Season
}
