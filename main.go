package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"suno-api/api"
	_ "suno-api/docs"
)

// @title Suno AI API
// @version v1.0
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// 文档地址
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.RegisterAPIHandler(r)

	r.Run(":3000")
}
