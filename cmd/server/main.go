package main

import (
	"fmt"
	"gin/internal/common"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := common.LoadConfig()

	common.InitDatabase(cfg)
	common.InitRedis(cfg)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
