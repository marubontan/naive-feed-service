package main

import (
	"log/slog"
	"naive-feed-service/app/config"
	infrastructure "naive-feed-service/app/infrastructure/repository"
	"naive-feed-service/app/server"
	"naive-feed-service/app/util"
	"os"
)

func main() {
	config, err := config.Load()
	if err != nil {
		util.Logger.Error("Failed to load config", slog.Any("err", err))
		os.Exit(1)
	}
	util.Logger.Info("Loaded config")
	db := infrastructure.NewDb(&config.Db)
	db.AutoMigrate(&infrastructure.FeedItem{})

	repositories := server.NewRepositories(db)
	ginServer := server.NewServer(repositories)
	ginServer.Setup()
	ginServer.Run(config.Server)
}
