package testing

import (
	"bytes"
	"clinic/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDoctor(t *testing.T) {
	router := gin.Default()
	payload := []byte(`{
		"name": "Elda",
    "gender": "female",
    "address": "bandung",
    "phone": "0822172",
    "specialist": "surgeon"
	}`)
	router.POST("/api/v1/doctor", handlers.CreateDoctor)
	req, err := http.NewRequest("POST", "/api/v1/doctor", bytes.NewReader(payload))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestGetDoctor(t *testing.T) {
	router := gin.Default()
	router.GET("/api/v1/doctor", handlers.GetDoctor)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/doctor", nil)

	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestUpdateDoctor(t *testing.T) {
	router := gin.Default()
	router.PUT("/api/v1/doctor/:id", handlers.UpdateDoctor)
	payload := []byte(`
	{
		"name": "dimasa",
		"gender": "male",
		"address": "bandung",
		"phone": "23232323",
		"specialist": "general"
	}`)
	req, err := http.NewRequest("PUT", "/api/v1/doctor/4", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
}

func TestDeleteDoctor(t *testing.T) {
	router := gin.Default()
	router.DELETE("/api/v1/doctor/:id", handlers.DeleteDoctor)
	req, err := http.NewRequest("DELETE", "/api/v1/doctor/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
