package usecase

import (
	entity "naive-feed-service/app/domain/entity"
	domain "naive-feed-service/app/domain/repository"
	"time"
)

type SaveFeedItemInputDTO struct {
	Id          string
	ItemId      string
	OrderNumber int
	CreatedAt   time.Time
}

type SaveFeedItemUsecase struct {
	feedRepository domain.FeedRepository
}

func NewSaveFeedItemUsecase(feedRepository domain.FeedRepository) *SaveFeedItemUsecase {
	return &SaveFeedItemUsecase{
		feedRepository: feedRepository,
	}
}

func (u *SaveFeedItemUsecase) Run(inputDTO *SaveFeedItemInputDTO) {
	u.feedRepository.Save(&entity.FeedItem{
		Id:          inputDTO.Id,
		ItemId:      inputDTO.ItemId,
		OrderNumber: inputDTO.OrderNumber,
		CreatedAt:   inputDTO.CreatedAt,
	})
}
