package usecase

import (
	entity "naive-feed-service/app/domain/entity"
	domain "naive-feed-service/app/domain/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetFeedUsecase(t *testing.T) {
	feedRepository := domain.NewMockFeedRepository(gomock.NewController(t))
	feedRepository.FeedTable["test"] = &entity.FeedItem{
		Id:          "test",
		ItemId:      "test",
		OrderNumber: 1,
		CreatedAt:   time.Now(),
	}
	useCase := NewGetFeedUsecase(feedRepository)
	feed := useCase.Run()
	assert.Equal(t, feed, []*entity.FeedItem{feedRepository.FeedTable["test"]})

}
