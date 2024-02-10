package infrastructure

import (
	"errors"
	"math/rand"
	"time"

	entity "naive-feed-service/app/domain/entity"

	"gorm.io/gorm"
)

type FeedItem struct {
	gorm.Model
	Id          string    `gorm:"column:id"`
	ItemId      string    `gorm:"column:item_id"`
	OrderNumber int       `gorm:"column:order_number"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

type FeedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *FeedRepository {
	return &FeedRepository{
		db: db,
	}
}

func (r *FeedRepository) Save(feed *entity.FeedItem) error {
	feedItem := FeedItem{
		Id:          feed.Id,
		ItemId:      feed.ItemId,
		OrderNumber: feed.OrderNumber,
		CreatedAt:   feed.CreatedAt,
	}
	result := r.db.Create(&feedItem)
	return result.Error
}

func (r *FeedRepository) GetAll() []*entity.FeedItem {
	var feedItems []FeedItem
	r.db.Find(&feedItems)
	var result []*entity.FeedItem
	for _, feedItem := range feedItems {
		result = append(result, &entity.FeedItem{
			Id:          feedItem.Id,
			ItemId:      feedItem.ItemId,
			OrderNumber: feedItem.OrderNumber,
			CreatedAt:   feedItem.CreatedAt,
		})
	}
	return result

}

func (r *FeedRepository) GetMinItemNumber() (int, error) {
	var minNumber int
	if err := r.db.Model(&FeedItem{}).Select("order_number").Order("order_number ASC").First(&minNumber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		} else {
			return -1, err
		}

	}
	return minNumber, nil

}

func (r *FeedRepository) Update() error {
	allFeedItems := r.GetAll()
	feedItemsOrder := make([]int, len(allFeedItems))
	for i := 0; i < len(allFeedItems); i++ {
		feedItemsOrder[i] = i
	}
	shuffleArray(feedItemsOrder)

	for i, feedItem := range allFeedItems {
		feedItem.OrderNumber = feedItemsOrder[i]
		r.db.Save(feedItem)
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
