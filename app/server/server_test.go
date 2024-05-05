//go:build integration

package server

import (
	"bytes"
	"encoding/json"
	"io"
	"naive-feed-service/app/domain/feed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*Server, *feed.MockFeedRepository) {
	feedRepository := feed.NewMockFeedRepository(gomock.NewController(t))
	repositories := Repositories{
		FeedRepository: feedRepository,
	}
	ginServer := NewServer(&repositories)
	ginServer.Setup()
	return ginServer, feedRepository

}

func TestHealth(t *testing.T) {
	ginServer, _ := setup(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	ginServer.engine.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

type Response struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

func TestPost(t *testing.T) {
	ginServer, feedRepository := setup(t)
	w := httptest.NewRecorder()
	requestBody := []byte(`{"item_id": "test"}`)
	req, _ := http.NewRequest("POST", "/feed", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	ginServer.engine.ServeHTTP(w, req)
	resBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error(err)
	}
	var res Response
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Received POST request", res.Message)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res.ID, feedRepository.FeedTable[res.ID].Id)

}
