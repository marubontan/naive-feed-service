package usecase

import (
	"naive-feed-service/app/domain/feed"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveFeedItemUsecase(t *testing.T) {
	feedRepository := feed.NewMockFeedRepository(gomock.NewController(t))
	uc := NewSaveFeedItemUsecase(feedRepository)
	feed := &feed.FeedItem{
		Id:          "test",
		ItemId:      "test",
		OrderNumber: -1,
		CreatedAt:   time.Now(),
	}
	id, err := uc.Run(&SaveFeedItemInputDTO{
		ItemId:    feed.ItemId,
		CreatedAt: feed.CreatedAt,
	})
	assert.Equal(t, nil, err)
	feed.Id = id
	assert.Equal(t, feed, feedRepository.FeedTable[feed.Id])

}
