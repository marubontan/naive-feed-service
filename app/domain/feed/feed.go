package feed

type FeedRepository interface {
	Save(feed *FeedItem) error
	Update() error
	GetAll() []*FeedItem
	GetMinItemNumber() (int, error)
}
