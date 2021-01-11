package services

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// AllCatalogues returns all items
func AllCatalogues() (*[]models.Catalogue, error) {
	var err error
	var db = utils.DBConn()

	catalogues := []models.Catalogue{}
	err = db.Model(&models.Catalogue{}).Find(&catalogues).Error

	if err != nil {
		fmt.Println(err)
		return &[]models.Catalogue{}, err
	}

	return &catalogues, nil
}

// GetCatalogue returns an item
func GetCatalogue(id int) (*models.Catalogue, error) {
	var catalogue models.Catalogue
	var err error
	var db = utils.DBConn()

	err = db.Find(&catalogue, id).Error

	if err != nil {
		return &models.Catalogue{}, err
	}

	return &catalogue, nil
}

// SaveCatalogue creates a new item
func SaveCatalogue(catalogue *models.Catalogue) (*models.Catalogue, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Catalogue{}).Create(&catalogue).Error; err != nil {
		return &models.Catalogue{}, err
	}

	return catalogue, nil
}

// UpdateCatalogue updates the item
func UpdateCatalogue(catalogue *models.Catalogue, id int) (*models.Catalogue, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Catalogue{}).Where("id = ?", id).Updates(&catalogue).Error; err != nil {
		return &models.Catalogue{}, err
	}

	return catalogue, nil
}

// DeleteCatalogue removes the item
func DeleteCatalogue(id int) (*models.Catalogue, error) {
	var err error
	var catalogue models.Catalogue
	var db = utils.DBConn()

	err = db.Find(&catalogue, id).Error

	if err != nil {
		return &models.Catalogue{}, err
	}

	err = db.Delete(&catalogue, id).Error

	if err != nil {
		return &models.Catalogue{}, err
	}

	return &catalogue, nil
}
