package models

import "github.com/jinzhu/gorm"

// Catalogue struct
type Catalogue struct {
	gorm.Model
	// ID          uint     `json:"id" gorm:"primaryKey"`
	Title       string   `json:"title" gorm:"unique;not null"`
	Slug        string   `json:"slug" gorm:"unique;not null"`
	Description string   `json:"description" gorm:"not null"`
	Price       float64  `json:"price" gorm:"not null"`
	MediaID     int      `json:"media_id"`
	OwnerID     int      `json:"owner_id"`
	Status      int      `json:"status" sql:"DEFAULT:1"`
	Seasons     []Season `json:"seasons"`
}
