package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/gnosi-api/app/models"
	seasonService "github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
	"github.com/gorilla/mux"
)

// GetSeasons get an item
func GetSeasons(w http.ResponseWriter, r *http.Request) {
	seasons, err := seasonService.AllSeasons()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, seasons)
}

// GetSeason gets an item
func GetSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	season, err := seasonService.GetSeason(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, season)
}

// CreateSeason creates an item
func CreateSeason(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	season := models.Season{}
	err = json.Unmarshal(body, &season)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	season.Slug = utils.Slugfy(season.Title)
	newCategory, err := seasonService.SaveSeason(&season)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newCategory)
}

// UpdateSeason updates an item
func UpdateSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	season := models.Season{}
	err = json.Unmarshal(body, &season)

	season.Slug = utils.Slugfy(season.Title)
	updatedCategory, err := seasonService.UpdateSeason(&season, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedCategory)
}

// DeleteSeason removes an item
func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = seasonService.DeleteSeason(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
