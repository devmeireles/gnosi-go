package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/gnosi-api/app/utils/validations"

	"github.com/devmeireles/gnosi-api/app/models"
	categoryService "github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
	"github.com/gorilla/mux"
)

// GetCategories get an item
func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := categoryService.AllCategories()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, categories)
}

// GetCategory gets an item
func GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	category, err := categoryService.GetCategory(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, category)
}

// CreateCategory creates an item
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	category := models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	category.Slug = utils.Slugfy(category.Title)

	validation := validations.ValidateCategory(&category)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	newCategory, err := categoryService.SaveCategory(&category)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newCategory)
}

// UpdateCategory updates an item
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = categoryService.GetCategory(id)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	category := models.Category{}
	err = json.Unmarshal(body, &category)

	category.Slug = utils.Slugfy(category.Title)

	validation := validations.ValidateCategory(&category)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	updatedCategory, err := categoryService.UpdateCategory(&category, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedCategory)
}

// DeleteCategory removes an item
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = categoryService.DeleteCategory(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
