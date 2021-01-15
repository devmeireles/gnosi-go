package models

import "github.com/jinzhu/gorm"

// Episode struct
type Episode struct {
	gorm.Model
	SeasonID    uint   `json:"season_id" gorm:"foreignkey:ID"`
	Title       string `json:"title" gorm:"unique;not null"`
	MediaID     uint   `json:"media_id" gorm:"not null"`
	Description string `json:"description"`
}
