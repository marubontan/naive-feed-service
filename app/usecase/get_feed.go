package usecase

import (
	"naive-feed-service/app/domain/feed"
)

type GetFeedUsecase struct {
	feedRepository feed.FeedRepository
}

func NewGetFeedUsecase(feedRepository feed.FeedRepository) *GetFeedUsecase {
	return &GetFeedUsecase{
		feedRepository: feedRepository,
	}
}

func (u *GetFeedUsecase) Run() []*feed.FeedItem {
	feed := u.feedRepository.GetAll()
	return feed
}
