// @title secSafe API
// @version 1.0
// @description This is the API documentation for secSafe
// @termsOfService http://swagger.io/terms/

// @contact.name John Kiribata
// @contact.url https://github.com/thronecoder/secsafe-backend
// @contact.email johnkiribata12@gmail.com

// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"

	"github.com/thronecoder/secsafe-backend/internal/common"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/thronecoder/secsafe-backend/api/docs"
)

// PingExample godoc
// @Summary      Ping the server
// @Description  Check if server is alive
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	cfg := common.LoadConfig()

	common.InitDatabase(cfg)
	common.InitRedis(cfg)

	r := gin.Default()

	//Register Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingHandler)

	r.Run(fmt.Sprintf(":%s", cfg.AppPort))
}
