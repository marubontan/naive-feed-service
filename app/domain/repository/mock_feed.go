package domain

import (
	"math/rand"
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

func (m *MockFeedRepository) Save(feed *entity.FeedItem) error {
	m.FeedTable[feed.Id] = feed
	return nil
}

func (m *MockFeedRepository) GetAll() []*entity.FeedItem {
	var result []*entity.FeedItem
	for _, feed := range m.FeedTable {
		result = append(result, feed)
	}
	return result
}

func (m *MockFeedRepository) GetMinItemNumber() (int, error) {
	minNumber := 0
	for _, feed := range m.FeedTable {
		if feed.OrderNumber < minNumber {
			minNumber = feed.OrderNumber
		}
	}
	return minNumber, nil
}

func (m *MockFeedRepository) Update() error {
	feedItemsOrder := make([]int, len(m.FeedTable))
	for i := 0; i < len(m.FeedTable); i++ {
		feedItemsOrder[i] = i
	}
	shuffleArray(feedItemsOrder)

	keys := make([]string, 0, len(m.FeedTable))
	for key := range m.FeedTable {
		keys = append(keys, key)
	}

	for i, key := range keys {
		m.FeedTable[key].OrderNumber = feedItemsOrder[i]
	}

	return nil

}

func shuffleArray(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
