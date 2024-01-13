package infrastructure

import (
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

func (r *FeedRepository) Save(feed *entity.FeedItem) {
	feedItem := FeedItem{
		Id:          feed.Id,
		ItemId:      feed.ItemId,
		OrderNumber: feed.OrderNumber,
		CreatedAt:   feed.CreatedAt,
	}
	r.db.Create(&feedItem)
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

func (r *FeedRepository) GetMinItemNumber() int {
	var minNumber int
	// TODO: Error handling
	r.db.Model(&FeedItem{}).Select("MIN(order_number)").Row().Scan(&minNumber)
	return minNumber
}
