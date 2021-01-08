package services

import (
	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

func GetAllCategories() (*[]models.Category, error) {
	var err error
	var db = utils.DBConn()

	categories := []models.Category{}
	err = db.Model(&models.Category{}).Find(&categories).Error

	if err != nil {
		return &[]models.Category{}, err
	}

	return &categories, nil
}
