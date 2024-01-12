package usecase

import (
	entity "naive-feed-service/app/domain/entity"
	domain "naive-feed-service/app/domain/repository"
)

type GetFeedUsecase struct {
	feedRepository domain.FeedRepository
}

func NewGetFeedUsecase(feedRepository domain.FeedRepository) *GetFeedUsecase {
	return &GetFeedUsecase{
		feedRepository: feedRepository,
	}
}

func (u *GetFeedUsecase) Run() []*entity.FeedItem {
	feed := u.feedRepository.GetAll()
	return feed
}
