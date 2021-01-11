package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// Login validates and generates a token
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userLogin := models.UserLogin{}
	err = json.Unmarshal(body, &userLogin)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	userReceived, err := services.GetByUsername(userLogin.Username)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	if !utils.ComparePasswords(userReceived.Password, []byte(userLogin.Password)) {
		utils.ResErr(w, errors.New("user and password not match"), http.StatusForbidden)
		return
	}

	token, _ := utils.CreateToken(userLogin.Username)

	utils.ResSuc(w, map[string]interface{}{"token": token})
}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	user.Password = utils.HashAndSalt([]byte(user.Password))
	newUser, err := services.CreateUser(&user)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newUser)
}
