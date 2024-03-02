package usecase

import (
	"naive-feed-service/app/domain/feed"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetFeedUsecase(t *testing.T) {
	feedRepository := feed.NewMockFeedRepository(gomock.NewController(t))
	feedRepository.FeedTable["test"] = &feed.FeedItem{
		Id:          "test",
		ItemId:      "test",
		OrderNumber: 1,
		CreatedAt:   time.Now(),
	}
	uc := NewGetFeedUsecase(feedRepository)
	actualFeed := uc.Run()
	assert.Equal(t, actualFeed, []*feed.FeedItem{feedRepository.FeedTable["test"]})

}
