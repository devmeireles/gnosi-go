package models

import "gorm.io/gorm"

// Category struct
type Category struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique;not null"`
	Slug        string `json:"slug" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
	Status      int    `json:"status" sql:"DEFAULT:1"`
}
