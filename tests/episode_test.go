package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/stretchr/testify/assert"
)

func TestEpisodeModule(t *testing.T) {
	t.Run("It should create a episode", func(t *testing.T) {
		var episode = models.Episode{
			SeasonID:    1,
			Title:       "Redis",
			Description: "A short desc",
		}

		episodeSave, _ := json.Marshal(episode)

		req, _ := http.NewRequest("POST", "/api/episode", bytes.NewBuffer(episodeSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return all episodes", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/episode", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return a episode", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/episode/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't return a nonexistent episode", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/episode/1190347", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't return an unparseable episode", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/episode/1190347x", nil)
		response := ExecuteRequest(req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't create a episode because this episode already exists", func(t *testing.T) {
		var episode = models.Episode{
			SeasonID:    1,
			Title:       "Redis",
			Description: "A short desc",
		}

		episodeSave, _ := json.Marshal(episode)

		req, _ := http.NewRequest("POST", "/api/episode", bytes.NewBuffer(episodeSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("It should update a episode", func(t *testing.T) {
		var episode = models.Episode{
			SeasonID:    1,
			Title:       "Redis",
			Description: "A short desc",
		}

		episodeSave, _ := json.Marshal(episode)

		req, _ := http.NewRequest("PUT", "/api/episode/1", bytes.NewBuffer(episodeSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't update a episode because this episode doesn't exists", func(t *testing.T) {
		var episode = models.Episode{
			SeasonID:    1,
			Title:       "Redis",
			Description: "A short desc",
		}

		episodeSave, _ := json.Marshal(episode)

		req, _ := http.NewRequest("PUT", "/api/episode/1983204", bytes.NewBuffer(episodeSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't update a episode because the body is missing", func(t *testing.T) {
		var episode = models.Episode{}

		episodeSave, _ := json.Marshal(episode)

		req, _ := http.NewRequest("PUT", "/api/episode/1", bytes.NewBuffer(episodeSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("It shouldn't delete a episode because this doesnt exists", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/episode/5951688", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It should delete a episode", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/episode/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
