package server

import (
	"fmt"
	"log/slog"
	_ "naive-feed-service/app/cmd/docs"
	"naive-feed-service/app/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

type Server struct {
	config *config.Config
	engine *gin.Engine
	db     *gorm.DB
}

func NewServer() *Server {
	config, err := config.Load()
	if err != nil {
		util.Logger.Error("Failed to load config", slog.Any("err", err))
	}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", config.Db.Host, config.Db.User, config.Db.Password, config.Db.Port)), &gorm.Config{})
	if err != nil {
		util.Logger.Error("Failed to connect to database", slog.Any("err", err))
	}

	return &Server{
		config: config,
		engine: gin.Default(),
		db:     db,
	}
}

// @Summary Check health
// @Description Check health
// @Success 200
// @Router /health [get]
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}

// @Summary Update feed
// @Description Update feed
// @Success 200
// @Router /feed [put]
func (s *Server) updateFeedHandler(c *gin.Context) {
	feedRepository := infrastructure.NewFeedRepository(s.db)
	useCase := usecase.NewUpdateFeedUsecase(feedRepository)
	err := useCase.Run()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	responseData := gin.H{"message": "Received PUT request"}
	c.JSON(200, responseData)

}

// @Summary Post feed item
// @Description Update feed
// @Param data body SaveFeedRequest true "Item to add to feed"
// @Success 200
// @Router /feed [post]
func (s *Server) postFeedItemHandler(c *gin.Context) {
	var requestData SaveFeedRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	feedRepository := infrastructure.NewFeedRepository(s.db)
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
}

// @Summary Get feed
// @Description get feed
// @Success 200
// @Router /feed [get]
func (s *Server) getFeedHandler(c *gin.Context) {
	feedRepository := infrastructure.NewFeedRepository(s.db)
	useCase := usecase.NewGetFeedUsecase(feedRepository)
	feedItems := useCase.Run()
	c.JSON(200, feedItems)
}

// @title Naive Feed Service
func (s *Server) Run() {
	util.Logger.Info("Starting server...")
	s.db.AutoMigrate(&infrastructure.FeedItem{})
	s.engine.GET("/health", s.healthHandler)
	s.engine.POST("/feed", s.postFeedItemHandler)
	s.engine.GET("/feed", s.getFeedHandler)
	s.engine.PUT("/feed", s.updateFeedHandler)
	s.engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.engine.Run(fmt.Sprint(s.config.Server.Address, ":", s.config.Server.Port))

}
