package usecase

import "naive-feed-service/app/domain/feed"

type UpdateFeedUsecase struct {
	feedRepository feed.FeedRepository
}

func NewUpdateFeedUsecase(feedRepository feed.FeedRepository) *UpdateFeedUsecase {
	return &UpdateFeedUsecase{
		feedRepository: feedRepository,
	}
}

func (u *UpdateFeedUsecase) Run() error {
	err := u.feedRepository.Update()
	return err

}
