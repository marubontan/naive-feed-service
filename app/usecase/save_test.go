package usecase

import (
	entity "naive-feed-service/app/domain/entity"
	domain "naive-feed-service/app/domain/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSaveFeedItemUsecase(t *testing.T) {
	feedRepository := domain.NewMockFeedRepository(gomock.NewController(t))
	useCase := NewSaveFeedItemUsecase(feedRepository)
	feed := &entity.FeedItem{
		Id:          "test",
		ItemId:      "test",
		OrderNumber: 1,
		CreatedAt:   time.Now(),
	}
	useCase.Run(&SaveFeedItemInputDTO{
		Id:          feed.Id,
		ItemId:      feed.ItemId,
		OrderNumber: feed.OrderNumber,
		CreatedAt:   feed.CreatedAt,
	})
	assert.Equal(t, feed, feedRepository.FeedTable[feed.Id])

}