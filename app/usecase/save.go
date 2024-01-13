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

func (u *SaveFeedItemUsecase) Run(inputDTO *SaveFeedItemInputDTO) string {
	id := uuid.New()
	presentMinOrderNumber := u.feedRepository.GetMinItemNumber()
	u.feedRepository.Save(&entity.FeedItem{
		Id:          id.String(),
		ItemId:      inputDTO.ItemId,
		OrderNumber: presentMinOrderNumber - 1,
		CreatedAt:   inputDTO.CreatedAt,
	})
	return id.String()
}
