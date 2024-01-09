package server

import (
	"fmt"
	"naive-feed-service/app/config"
	"naive-feed-service/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	util.Logger.Println("Starting server...")
	config, err := config.Load()
	if err != nil {
		util.Logger.Fatal(err)
	}
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(fmt.Sprint(config.Server.Address, ":", config.Server.Port))

}
