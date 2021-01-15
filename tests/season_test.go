package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSeasons(t *testing.T) {
	t.Run("It should create a season", func(t *testing.T) {
		var season = models.Season{
			CatalogueID: 1,
			Title:       "Redis",
			Description: "A short desc",
		}

		seasonSave, _ := json.Marshal(season)

		req, _ := http.NewRequest("POST", "/api/season", bytes.NewBuffer(seasonSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return all seasons", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/season", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return a season", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/season/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't return a nonexistent season", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/season/1190347", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't return an unparseable season", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/season/1190347x", nil)
		response := ExecuteRequest(req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't create a season because this season already exists", func(t *testing.T) {
		var season = models.Season{
			CatalogueID: 1,
			Title:       "Redis",
			Description: "A short desc",
		}

		seasonSave, _ := json.Marshal(season)

		req, _ := http.NewRequest("POST", "/api/season", bytes.NewBuffer(seasonSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("It should update a season", func(t *testing.T) {
		var season = models.Season{
			CatalogueID: 1,
			Title:       "Redis",
			Description: "A short desc",
		}

		seasonSave, _ := json.Marshal(season)

		req, _ := http.NewRequest("PUT", "/api/season/1", bytes.NewBuffer(seasonSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't update a season because this season doesn't exists", func(t *testing.T) {
		var season = models.Season{
			CatalogueID: 1,
			Title:       "Redis",
			Description: "A short desc",
		}

		seasonSave, _ := json.Marshal(season)

		req, _ := http.NewRequest("PUT", "/api/season/1983204", bytes.NewBuffer(seasonSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't update a season because the body is missing", func(t *testing.T) {
		var season = models.Season{}

		seasonSave, _ := json.Marshal(season)

		req, _ := http.NewRequest("PUT", "/api/season/1", bytes.NewBuffer(seasonSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("It shouldn't delete a season because this doesnt exists", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/season/5951688", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It should delete a season", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/season/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	e := os.Remove("database/gorm.db")
	if e != nil {
		log.Fatal(e)

	}
}
