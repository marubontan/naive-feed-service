package usecase

import (
	"naive-feed-service/app/domain/feed"
	"time"

	"github.com/google/uuid"
)

type SaveFeedItemInputDTO struct {
	ItemId    string
	CreatedAt time.Time
}

type SaveFeedItemUsecase struct {
	feedRepository feed.FeedRepository
}

func NewSaveFeedItemUsecase(feedRepository feed.FeedRepository) *SaveFeedItemUsecase {
	return &SaveFeedItemUsecase{
		feedRepository: feedRepository,
	}
}

func (u *SaveFeedItemUsecase) Run(inputDTO *SaveFeedItemInputDTO) (string, error) {
	id := uuid.New()
	presentMinOrderNumber, err := u.feedRepository.GetMinItemNumber()
	if err != nil {
		return "", err
	}
	err = u.feedRepository.Save(&feed.FeedItem{
		Id:          id.String(),
		ItemId:      inputDTO.ItemId,
		OrderNumber: presentMinOrderNumber - 1,
		CreatedAt:   inputDTO.CreatedAt,
	})
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
