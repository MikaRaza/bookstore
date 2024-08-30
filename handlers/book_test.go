package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example/bookstore/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	err = database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func TestGetBooks(t *testing.T) {
	setup()

	router := gin.Default()
	router.GET("/books", GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Vous pouvez ajouter des vérifications supplémentaires ici
}
