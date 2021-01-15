package services

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// AllEpisodes returns all items
func AllEpisodes() (*[]models.Episode, error) {
	var err error
	var db = utils.DBConn()

	episodes := []models.Episode{}
	err = db.Model(&models.Episode{}).Find(&episodes).Error

	if err != nil {
		return nil, err
	}

	return &episodes, nil
}

// GetEpisode returns an item
func GetEpisode(id int) (*models.Episode, error) {
	var episode models.Episode
	var err error
	var db = utils.DBConn()

	err = db.First(&episode, id).Error

	if err != nil {
		return nil, err
	}

	return &episode, nil
}

// SaveEpisode creates a new item
func SaveEpisode(episode *models.Episode) (*models.Episode, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Episode{}).Create(&episode).Error; err != nil {
		return nil, err
	}

	return episode, nil
}

// UpdateEpisode updates the item
func UpdateEpisode(episode *models.Episode, id int) (*models.Episode, error) {
	var db = utils.DBConn()

	err := db.Model(&models.Episode{}).Where("id = ?", id).Updates(&episode).Error
	// err := db.Save(episode).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return episode, nil
}

// DeleteEpisode removes the item
func DeleteEpisode(id int) (*models.Episode, error) {
	var err error
	var episode models.Episode
	var db = utils.DBConn()

	err = db.First(&episode, id).Error

	if err != nil {
		return nil, err
	}

	err = db.Delete(&episode, id).Error

	if err != nil {
		return nil, err
	}

	return &episode, nil
}
