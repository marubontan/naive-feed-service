//go:build integration

package infrastructure

import (
	"database/sql"
	"fmt"
	"log/slog"
	"naive-feed-service/app/domain/feed"
	"naive-feed-service/app/util"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var databaseUrl string

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		slog.Error("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		slog.Error("Could not connect to Docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_USER=postgres",
			"POSTGRES_DB=postgres",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		slog.Error("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl = fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", hostAndPort)

	resource.Expire(120)

	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		slog.Error("Could not connect to docker:", err)
	}
	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		slog.Error("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func newMockDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		util.Logger.Error("Failed to connect to database", slog.Any("err", err))
		os.Exit(1)
	}
	return db
}

func TestRepositorySave(t *testing.T) {
	mockDb := newMockDb()
	mockDb.AutoMigrate(&FeedItem{})
	r := NewFeedRepository(mockDb)
	feed := &feed.FeedItem{
		Id:          "1",
		ItemId:      "2",
		OrderNumber: 3,
		CreatedAt:   time.Now(),
	}
	r.Save(feed)
	var savedFeedItem FeedItem
	mockDb.First(&savedFeedItem, "id=?", feed.Id)
	assert.Equal(t, feed.Id, savedFeedItem.Id)
	assert.Equal(t, feed.ItemId, savedFeedItem.ItemId)
	assert.Equal(t, feed.OrderNumber, savedFeedItem.OrderNumber)
}
