package articles_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/2golang/vnexpress-rss/handlers/articles"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type response struct {
	Message string `json:"message,omitempty"`
	ID      string `json:"id,omitempty"`
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

func TestListArticles(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedCode   int
		expectedResult response
	}{
		{
			name:           "Success GET request",
			method:         http.MethodGet,
			expectedCode:   http.StatusOK,
			expectedResult: response{Message: "pong pong"},
		},
		{
			name:           "Wrong Method",
			method:         http.MethodPost,
			expectedCode:   http.StatusMethodNotAllowed,
			expectedResult: response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupTestRouter()
			router.GET("/articles", articles.ListArticles)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, "/articles", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusOK {
				var got response
				err := json.Unmarshal(w.Body.Bytes(), &got)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, got)
			}
		})
	}
}

func TestGetArticleById(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		method         string
		expectedCode   int
		expectedResult response
	}{
		{
			name:           "Valid ID",
			id:             "123",
			method:         http.MethodGet,
			expectedCode:   http.StatusOK,
			expectedResult: response{ID: "123"},
		},
		{
			name:           "Empty ID",
			id:             "",
			method:         http.MethodGet,
			expectedCode:   http.StatusNotFound,
			expectedResult: response{},
		},
		{
			name:           "Special Characters ID",
			id:             "123@#$",
			method:         http.MethodGet,
			expectedCode:   http.StatusOK,
			expectedResult: response{ID: "123@#$"},
		},
		{
			name:           "Wrong Method",
			id:             "123",
			method:         http.MethodPost,
			expectedCode:   http.StatusMethodNotAllowed,
			expectedResult: response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupTestRouter()
			router.GET("/articles/:id", articles.GetArticleById)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, "/articles/"+tt.id, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusOK {
				var got response
				err := json.Unmarshal(w.Body.Bytes(), &got)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, got)
			}
		})
	}
}
