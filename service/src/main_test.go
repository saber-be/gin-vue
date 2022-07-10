package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"qgen/challange/controllers"
	"qgen/challange/databases"
	"qgen/challange/middlewares"
	"qgen/challange/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	DBAdapter := databases.NewMOCKAdapter()
	router.Use(middlewares.DBMiddleware(DBAdapter))
	return router
}

func TestUserAgeCriteria(t *testing.T) {
	r := SetUpRouter()
	r.POST("/", controllers.CreateUser)

	user := models.UserInput{
		Name:  "saber",
		Email: "test@gmail.com",
		Phone: "123-123-1234",
		Age:   17,
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := ioutil.ReadAll(w.Body)
	mockResponse := `{"error":"Key: 'UserInput.Age' Error:Field validation for 'Age' failed on the 'min' tag"}`
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Test for age >= 18
	user.Age = 18
	jsonValue, _ = json.Marshal(user)
	req, _ = http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
