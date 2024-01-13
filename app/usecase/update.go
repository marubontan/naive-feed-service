package usecase

import domain "naive-feed-service/app/domain/repository"

type UpdateFeedUsecase struct {
	feedRepository domain.FeedRepository
}

func NewUpdateFeedUsecase(feedRepository domain.FeedRepository) *UpdateFeedUsecase {
	return &UpdateFeedUsecase{
		feedRepository: feedRepository,
	}
}

func (u *UpdateFeedUsecase) Run() error {
	err := u.feedRepository.Update()
	return err

}
