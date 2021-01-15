package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/gnosi-api/app/models"
	catalogueService "github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
	"github.com/devmeireles/gnosi-api/app/utils/validations"
	"github.com/gorilla/mux"
)

// GetCatalogues get an item
func GetCatalogues(w http.ResponseWriter, r *http.Request) {
	catalogues, err := catalogueService.AllCatalogues()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, catalogues)
}

// GetCatalogue gets an item
func GetCatalogue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	catalogue, err := catalogueService.GetCatalogue(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, catalogue)
}

// CreateCatalogue creates an item
func CreateCatalogue(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	catalogue := models.Catalogue{}
	err = json.Unmarshal(body, &catalogue)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	validation := validations.ValidateCatalogue(&catalogue)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	catalogue.Slug = utils.Slugfy(catalogue.Title)
	newCatalogue, err := catalogueService.SaveCatalogue(&catalogue)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newCatalogue)
}

// UpdateCatalogue updates an item
func UpdateCatalogue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = catalogueService.GetCatalogue(id)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	catalogue := models.Catalogue{}
	err = json.Unmarshal(body, &catalogue)

	catalogue.Slug = utils.Slugfy(catalogue.Title)

	validation := validations.ValidateCatalogue(&catalogue)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	updatedCatalogue, err := catalogueService.UpdateCatalogue(&catalogue, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedCatalogue)
}

// DeleteCatalogue removes an item
func DeleteCatalogue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = catalogueService.DeleteCatalogue(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
