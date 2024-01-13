package usecase

import (
	entity "naive-feed-service/app/domain/entity"
	domain "naive-feed-service/app/domain/repository"
	"time"

	"github.com/google/uuid"
)

type SaveFeedItemInputDTO struct {
	ItemId    string
	CreatedAt time.Time
}

type SaveFeedItemUsecase struct {
	feedRepository domain.FeedRepository
}

func NewSaveFeedItemUsecase(feedRepository domain.FeedRepository) *SaveFeedItemUsecase {
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
	err = u.feedRepository.Save(&entity.FeedItem{
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
