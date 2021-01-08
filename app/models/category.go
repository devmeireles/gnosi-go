package models

import "github.com/jinzhu/gorm"

// Category struct
type Category struct {
	gorm.Model
	Title       string `json:"title" gorm:"unique;not null"`
	Slug        string `json:"slug" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
	Status      int    `json:"status" sql:"DEFAULT:1"`
	CreatedAt   string `json:"created_at" gorm:"not null"`
	UpdatedAt   string `json:"updated_at" gorm:"not null"`
	DeletedAt   string `json:"deleted_at" gorm:"not null"`
}
