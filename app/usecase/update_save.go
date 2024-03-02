package usecase

import (
	"naive-feed-service/app/domain/feed"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateFeedUsecase(t *testing.T) {
	feedRepository := feed.NewMockFeedRepository(gomock.NewController(t))
	uc := NewUpdateFeedUsecase(feedRepository)
	err := uc.Run()
	assert.Equal(t, err, nil)
}
