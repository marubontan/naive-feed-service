package feed

import "time"

type FeedItem struct {
	Id          string
	ItemId      string
	OrderNumber int
	CreatedAt   time.Time
}
