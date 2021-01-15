package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/stretchr/testify/assert"
)

func TestCatalogueModule(t *testing.T) {
	t.Run("It should create a catalogue", func(t *testing.T) {
		var catalogue = models.Catalogue{
			Title:       "Redis",
			Description: "A short desc",
			Price:       10.9,
			MediaID:     1,
			OwnerID:     1,
		}

		catalogueSave, _ := json.Marshal(catalogue)

		req, _ := http.NewRequest("POST", "/api/catalogue", bytes.NewBuffer(catalogueSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return all catalogues", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/catalogue", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return a catalogue", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/catalogue/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't return a nonexistent catalogue", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/catalogue/1190347", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't return an unparseable catalogue", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/catalogue/1190347x", nil)
		response := ExecuteRequest(req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't create a catalogue because this catalogue already exists", func(t *testing.T) {
		var catalogue = models.Catalogue{
			Title:       "Redis",
			Description: "A short desc",
			Price:       10.9,
			MediaID:     1,
			OwnerID:     1,
		}

		catalogueSave, _ := json.Marshal(catalogue)

		req, _ := http.NewRequest("POST", "/api/catalogue", bytes.NewBuffer(catalogueSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("It should update a catalogue", func(t *testing.T) {
		var catalogue = models.Catalogue{
			Title:       "MySQL Course",
			Description: "A short desc for mysql",
			Price:       10.9,
			MediaID:     1,
			OwnerID:     1,
		}

		catalogueSave, _ := json.Marshal(catalogue)

		req, _ := http.NewRequest("PUT", "/api/catalogue/1", bytes.NewBuffer(catalogueSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't update a catalogue because this catalogue doesn't exists", func(t *testing.T) {
		var catalogue = models.Catalogue{
			Title:       "MySQL Course",
			Description: "A short desc for mysql",
		}

		catalogueSave, _ := json.Marshal(catalogue)

		req, _ := http.NewRequest("PUT", "/api/catalogue/1983204", bytes.NewBuffer(catalogueSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't update a catalogue because the body is missing", func(t *testing.T) {
		var catalogue = models.Catalogue{}

		catalogueSave, _ := json.Marshal(catalogue)

		req, _ := http.NewRequest("PUT", "/api/catalogue/1", bytes.NewBuffer(catalogueSave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("It shouldn't delete a catalogue because this doesnt exists", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/catalogue/5951688", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It should delete a catalogue", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/catalogue/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
