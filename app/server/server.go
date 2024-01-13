package server

import (
	"fmt"
	"naive-feed-service/app/config"
	infrastructure "naive-feed-service/app/infrastructure/repository"
	"naive-feed-service/app/usecase"
	"naive-feed-service/app/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SaveFeedRequest struct {
	ItemId string `json:"item_id"`
}

func Run() {
	util.Logger.Println("Starting server...")
	config, err := config.Load()
	if err != nil {
		util.Logger.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", config.Db.Host, config.Db.User, config.Db.Password, config.Db.Port)), &gorm.Config{})
	db.AutoMigrate(&infrastructure.FeedItem{})
	if err != nil {
		util.Logger.Fatal(err)
	}
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	engine.POST("/feed", func(c *gin.Context) {
		var requestData SaveFeedRequest

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		feedRepository := infrastructure.NewFeedRepository(db)
		useCase := usecase.NewSaveFeedItemUsecase(feedRepository)
		id, err := useCase.Run(&usecase.SaveFeedItemInputDTO{
			ItemId:    requestData.ItemId,
			CreatedAt: time.Now(),
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		responseData := gin.H{"message": "Received POST request", "id": id}
		c.JSON(200, responseData)

	})

	engine.GET("/feed", func(c *gin.Context) {
		feedRepository := infrastructure.NewFeedRepository(db)
		useCase := usecase.NewGetFeedUsecase(feedRepository)
		feedItems := useCase.Run()
		c.JSON(200, feedItems)
	})

	engine.PUT("/feed", func(c *gin.Context) {
		feedRepository := infrastructure.NewFeedRepository(db)
		useCase := usecase.NewUpdateFeedUsecase(feedRepository)
		err := useCase.Run()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		responseData := gin.H{"message": "Received PUT request"}
		c.JSON(200, responseData)
	})
	engine.Run(fmt.Sprint(config.Server.Address, ":", config.Server.Port))

}
