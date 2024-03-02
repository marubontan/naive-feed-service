package server

import (
	"fmt"
	"log/slog"
	_ "naive-feed-service/app/cmd/docs"
	"naive-feed-service/app/config"
	"naive-feed-service/app/domain/feed"
	infrastructure "naive-feed-service/app/infrastructure/repository"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"naive-feed-service/app/usecase"
	"naive-feed-service/app/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SaveFeedRequest struct {
	ItemId string `json:"item_id" binding:"required"`
}

type Repositories struct {
	FeedRepository feed.FeedRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		FeedRepository: infrastructure.NewFeedRepository(db),
	}
}

type Server struct {
	engine       *gin.Engine
	repositories Repositories
}

func NewServer(repositories Repositories) *Server {
	return &Server{
		engine:       gin.Default(),
		repositories: repositories,
	}
}

// @Summary Check health
// @Description Check health
// @Success 200
// @Router /health [get]
func (s *Server) healthHandler(c *gin.Context) {
	util.Logger.Info("Health check called")
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

// @Summary Update feed
// @Description Update feed
// @Success 200
// @Router /feed [put]
func (s *Server) updateFeedHandler(c *gin.Context) {
	util.Logger.Info("Update feed called")
	useCase := usecase.NewUpdateFeedUsecase(s.repositories.FeedRepository)
	err := useCase.Run()
	if err != nil {
		util.Logger.Info("Failed to update feed", slog.Any("err", err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	responseData := gin.H{"message": "Received PUT request"}
	c.JSON(200, responseData)
	util.Logger.Info("Updated feed")

}

// @Summary Post feed item
// @Description Update feed
// @Param data body SaveFeedRequest true "Item to add to feed"
// @Success 200
// @Router /feed [post]
func (s *Server) postFeedItemHandler(c *gin.Context) {
	util.Logger.Info("Post feed item called")
	var requestData SaveFeedRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.Logger.Error("Failed to bind JSON", slog.Any("err", err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	useCase := usecase.NewSaveFeedItemUsecase(s.repositories.FeedRepository)
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
	util.Logger.Info("Posted feed item")
}

// @Summary Get feed
// @Description get feed
// @Success 200
// @Router /feed [get]
func (s *Server) getFeedHandler(c *gin.Context) {
	util.Logger.Info("Get feed called")
	useCase := usecase.NewGetFeedUsecase(s.repositories.FeedRepository)
	feedItems := useCase.Run()
	c.JSON(200, feedItems)
	util.Logger.Info("Got feed")
}

// @title Naive Feed Service
func (s *Server) Setup() *gin.Engine {
	util.Logger.Info("Starting server...")
	s.engine.GET("/health", s.healthHandler)
	s.engine.POST("/feed", s.postFeedItemHandler)
	s.engine.GET("/feed", s.getFeedHandler)
	s.engine.PUT("/feed", s.updateFeedHandler)
	s.engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return s.engine
}

func (s *Server) Run(server_config config.Server) {
	s.engine.Run(fmt.Sprint(server_config.Address, ":", server_config.Port))
}
