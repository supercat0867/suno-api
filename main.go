package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"suno-api/api"
	_ "suno-api/docs"
)

// @title Suno AI API
// @version v1.0
func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 文档地址
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.RegisterAPIHandler(r)

	r.Run(":3000")
}
