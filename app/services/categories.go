package services

import (
	"fmt"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// AllCategories returns all items
func AllCategories() (*[]models.Category, error) {
	var err error
	var db = utils.DBConn()

	categories := []models.Category{}
	err = db.Model(&models.Category{}).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return &categories, nil
}

// GetCategory returns an item
func GetCategory(id int) (*models.Category, error) {
	var category models.Category
	var err error
	var db = utils.DBConn()

	err = db.First(&category, id).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// SaveCategory creates a new item
func SaveCategory(category *models.Category) (*models.Category, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.Category{}).Create(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

// UpdateCategory updates the item
func UpdateCategory(category *models.Category, id int) (*models.Category, error) {
	var db = utils.DBConn()

	err := db.Model(&models.Category{}).Where("id = ?", id).Updates(&category).Error
	// err := db.Save(category).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return category, nil
}

// DeleteCategory removes the item
func DeleteCategory(id int) (*models.Category, error) {
	var err error
	var category models.Category
	var db = utils.DBConn()

	err = db.Find(&category, id).Error

	if err != nil {
		return nil, err
	}

	err = db.Delete(&category, id).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}
