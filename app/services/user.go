package services

import (
	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/utils"
)

func CreateUser(user *models.User) (*models.User, error) {
	var db = utils.DBConn()

	if err := db.Model(&models.User{}).Create(&user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func GetAllUsers() (*[]models.User, error) {
	var db = utils.DBConn()
	var err error
	users := []models.User{}

	err = db.Model(&models.User{}).Find(&users).Error

	if err != nil {
		return &[]models.User{}, err
	}

	return &users, nil
}

func GetUser(id int) (*models.User, error) {
	var db = utils.DBConn()
	var user models.User
	var err error

	err = db.Where("id = ?", id).Preload("Address").Find(&user).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func GetByUsername(username string) (*models.User, error) {
	var db = utils.DBConn()
	var user models.User
	var err error

	err = db.Where("username = ?", username).Find(&user).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func DeleteUser(id int) (*models.User, error) {
	var err error
	var user models.User
	var db = utils.DBConn()

	err = db.Find(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	err = db.Delete(&user, id).Error

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}
