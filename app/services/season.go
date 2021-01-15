package services

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// AllSeasons returns all items
func AllSeasons() (*[]models.Season, error) {
	var err error
	var db = utils.DBConn()

	seasons := []models.Season{}
	err = db.Preload("Episodes").Model(&models.Season{}).Find(&seasons).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &seasons, nil
}

// GetSeason returns an item
func GetSeason(id int) (*models.Season, error) {
	var season models.Season
	var err error
	var db = utils.DBConn()

	err = db.Preload("Episodes").First(&season, id).Error

	if err != nil {
		return nil, err
	}

	return &season, nil
}

// SaveSeason creates a new item
func SaveSeason(season *models.Season) (*models.Season, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Season{}).Create(&season).Error; err != nil {
		return nil, err
	}

	return season, nil
}

// UpdateSeason updates the item
func UpdateSeason(season *models.Season, id int) (*models.Season, error) {
	var db = utils.DBConn()

	err := db.Model(&models.Season{}).Where("id = ?", id).Updates(&season).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return season, nil
}

// DeleteSeason removes the item
func DeleteSeason(id int) (*models.Season, error) {
	var err error
	var season models.Season
	var db = utils.DBConn()

	err = db.First(&season, id).Error

	if err != nil {
		return &models.Season{}, err
	}

	err = db.Delete(&season, id).Error

	if err != nil {
		return &models.Season{}, err
	}

	return &season, nil
}
