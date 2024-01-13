package usecase

import (
	domain "naive-feed-service/app/domain/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateFeedUsecase(t *testing.T) {
	feedRepository := domain.NewMockFeedRepository(gomock.NewController(t))
	useCase := NewUpdateFeedUsecase(feedRepository)
	err := useCase.Run()
	assert.Equal(t, err, nil)
}
