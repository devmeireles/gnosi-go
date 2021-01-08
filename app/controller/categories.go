package controller

import (
	"net/http"

	"github.com/devmeireles/gnosi-api/app/services"
	"github.com/devmeireles/gnosi-api/app/utils"
)

// GetAllCategories godoc
// @Summary Get details of all skills
// @Description Get details of all skills
// @Tags skills
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} models.Skill
// @Router /api/skills [get]
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	skills, err := services.GetAllCategories()

	if err != nil {
		utils.ResErr(w, err, http.StatusInternalServerError)
		return
	}

	utils.ResSuc(w, skills)
}
