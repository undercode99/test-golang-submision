package service_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"server-grpc/rest-api/handler"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestSearchMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := handler.SetupRouters()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/search-movie?keyword=Batman", nil)
	router.ServeHTTP(w, req)

	type SearchMovieResponse struct {
		Response string `json:"Response"`
	}
	type SearchResponse struct {
		Search SearchMovieResponse `json:"Search"`
	}
	var parsed SearchResponse
	err := json.Unmarshal(w.Body.Bytes(), &parsed)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "True", parsed.Search.Response)
}

func TestDetailMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := handler.SetupRouters()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/detail-movie/tt0076759", nil)
	router.ServeHTTP(w, req)

	type DetailResponse struct {
		Response string `json:"Response"`
		ImdbID   string `json:"ImdbID"`
	}
	var parsed DetailResponse
	err := json.Unmarshal(w.Body.Bytes(), &parsed)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "True", parsed.Response)
	assert.Equal(t, "tt0076759", parsed.ImdbID)
}
