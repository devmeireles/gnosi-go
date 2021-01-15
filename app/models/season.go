package models

import "github.com/jinzhu/gorm"

// Season struct
type Season struct {
	gorm.Model
	CatalogueID uint   `json:"catalogue_id" gorm:"foreignkey:ID"`
	Title       string `json:"title" gorm:"unique;not null"`
	Slug        string `json:"slug" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
