package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-auth/infrastructure/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPingRoute(t *testing.T) {
	mockResponse := `{"message":"pong"}`
	r := SetUpRouter()
	r.GET("/ping", controllers.PingController{}.Ping)
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, mockResponse, w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
