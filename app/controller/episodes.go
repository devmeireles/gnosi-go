package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/devmeireles/gnosi-api/app/utils/validations"

	"github.com/devmeireles/gnosi-api/app/models"
	episodeService "github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
	"github.com/gorilla/mux"
)

// GetEpisodes get an item
func GetEpisodes(w http.ResponseWriter, r *http.Request) {
	categories, err := episodeService.AllEpisodes()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, categories)
}

// GetEpisode gets an item
func GetEpisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	episode, err := episodeService.GetEpisode(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}
	utils.ResSuc(w, episode)
}

// CreateEpisode creates an item
func CreateEpisode(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	episode := models.Episode{}
	err = json.Unmarshal(body, &episode)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	validation := validations.ValidateEpisode(&episode)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	newEpisode, err := episodeService.SaveEpisode(&episode)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}
	utils.ResSuc(w, newEpisode)
}

// UpdateEpisode updates an item
func UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = episodeService.GetEpisode(id)
	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	episode := models.Episode{}
	err = json.Unmarshal(body, &episode)

	validation := validations.ValidateEpisode(&episode)

	if validation != nil {
		utils.ResValidation(w, validation)
		return
	}

	updatedEpisode, err := episodeService.UpdateEpisode(&episode, id)

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, updatedEpisode)
}

// DeleteEpisode removes an item
func DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	_, err = episodeService.DeleteEpisode(id)

	if err != nil {
		utils.ResErr(w, err, http.StatusNotFound)
		return
	}

	utils.ResTrue(w)
}
