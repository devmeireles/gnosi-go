package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devmeireles/gnosi-api/app/models"
	"github.com/stretchr/testify/assert"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Router.ServeHTTP(rr, req)

	return rr
}

func parseBody(content *httptest.ResponseRecorder) models.ResponseMsg {
	parsedRes := models.ResponseMsg{}
	body, _ := ioutil.ReadAll(content.Body)
	json.Unmarshal(body, &parsedRes)

	return parsedRes

}

func TestGetCategories(t *testing.T) {
	t.Run("Get all categories", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/category", nil)
		response := executeRequest(req)
		parsedBody := parseBody(response)

		assert.True(t, parsedBody.Success)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
