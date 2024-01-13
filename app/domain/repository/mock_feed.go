package domain

import (
	entity "naive-feed-service/app/domain/entity"

	"go.uber.org/mock/gomock"
)

type MockFeedRepository struct {
	*gomock.Controller
	FeedTable map[string]*entity.FeedItem
}

func NewMockFeedRepository(ctrl *gomock.Controller) *MockFeedRepository {
	return &MockFeedRepository{ctrl, make(map[string]*entity.FeedItem)}
}

func (m *MockFeedRepository) Save(feed *entity.FeedItem) {
	m.FeedTable[feed.Id] = feed
}

func (m *MockFeedRepository) GetAll() []*entity.FeedItem {
	var result []*entity.FeedItem
	for _, feed := range m.FeedTable {
		result = append(result, feed)
	}
	return result
}

func (m *MockFeedRepository) GetMinItemNumber() int {
	minNumber := 0
	for _, feed := range m.FeedTable {
		if feed.OrderNumber < minNumber {
			minNumber = feed.OrderNumber
		}
	}
	return minNumber
}
