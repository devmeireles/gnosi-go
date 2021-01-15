package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/stretchr/testify/assert"
)

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Router.ServeHTTP(rr, req)

	return rr
}

func ParseBody(content *httptest.ResponseRecorder) models.ResponseMsg {
	parsedRes := models.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}

func TestGetCategories(t *testing.T) {
	t.Run("It should create a category", func(t *testing.T) {
		var category = models.Category{
			Title:       "Redis",
			Description: "A short desc",
		}

		categorySave, _ := json.Marshal(category)

		req, _ := http.NewRequest("POST", "/api/category", bytes.NewBuffer(categorySave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return all categories", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/category", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It should return a category", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/category/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't return a nonexistent category", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/category/1190347", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't return an unparseable category", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/category/1190347x", nil)
		response := ExecuteRequest(req)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't create a category because this category already exists", func(t *testing.T) {
		var category = models.Category{
			Title:       "Redis",
			Description: "A short desc",
		}

		categorySave, _ := json.Marshal(category)

		req, _ := http.NewRequest("POST", "/api/category", bytes.NewBuffer(categorySave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("It should update a category", func(t *testing.T) {
		var category = models.Category{
			Title:       "MySQL",
			Description: "A short desc for mysql",
		}

		categorySave, _ := json.Marshal(category)

		req, _ := http.NewRequest("PUT", "/api/category/1", bytes.NewBuffer(categorySave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("It shouldn't update a category because this category doesn't exists", func(t *testing.T) {
		var category = models.Category{
			Title:       "MySQL",
			Description: "A short desc for mysql",
		}

		categorySave, _ := json.Marshal(category)

		req, _ := http.NewRequest("PUT", "/api/category/1983204", bytes.NewBuffer(categorySave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It shouldn't update a category because the body is missing", func(t *testing.T) {
		var category = models.Category{}

		categorySave, _ := json.Marshal(category)

		req, _ := http.NewRequest("PUT", "/api/category/1", bytes.NewBuffer(categorySave))

		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("It shouldn't delete a category because this doesnt exists", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/category/5951688", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.False(t, parsedBody.Success)
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("It should delete a category", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/category/1", nil)
		response := ExecuteRequest(req)
		parsedBody := ParseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
