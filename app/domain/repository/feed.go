package domain

import (
	entity "naive-feed-service/app/domain/entity"
)

type FeedRepository interface {
	Save(feed *entity.FeedItem) error
	Update() error
	GetAll() []*entity.FeedItem
	GetMinItemNumber() (int, error)
}
