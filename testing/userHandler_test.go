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

func TestRegister(t *testing.T) {
	router := gin.Default()

	payload := []byte(`
	{"name": "dimas",
	"email": "dimas@gmail.com",
	"password": "1234asdf"}`)

	router.POST("/api/v1/register", handlers.Register)

	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestLogin(t *testing.T) {
	router := gin.Default()

	payload := []byte(`
	{"email": "dimas@gmail.com",
	"password": "1234asdf"}`)

	router.POST("/api/v1/login", handlers.Login)

	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}
