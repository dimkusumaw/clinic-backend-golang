package testing

import (
	"bytes"
	"clinic/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	router := gin.Default()
	payload := []byte(`
	{
		"name": "dimasa",
		"age": 21,
		"gender": "male",
		"address": "bandung",
		"phone": "23232323"
	}`)

	router.POST("/api/v1/patient", handlers.CreatePatient)
	req, _ := http.NewRequest("POST", "/api/v1/patient", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestGetPatient(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/patient", handlers.GetPatient)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/patient", nil)

	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestUpdatePatient(t *testing.T) {
	router := gin.Default()
	router.PUT("/api/v1/patient/:id", handlers.UpdatePatient)
	payload := []byte(`
	{
		"name": "dimasa",
		"age": 19,
		"gender": "male",
		"address": "bandung",
		"phone": "23232323"
	}`)
	req, err := http.NewRequest("PUT", "/api/v1/patient/1", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestDeletePatient(t *testing.T) {
	router := gin.Default()
	router.DELETE("/api/v1/patient/:id", handlers.DeletePatient)
	req, err := http.NewRequest("DELETE", "/api/v1/patient/6", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
