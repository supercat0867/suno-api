package api

import (
	"github.com/gin-gonic/gin"
	"suno-api/controller"
	"suno-api/kernel"
)

func RegisterAPIHandler(r *gin.Engine) {
	client := kernel.NenClient()
	contr := controller.NewController(client)

	group := r.Group("/api/v1")
	group.POST("/generate/song", contr.GenerateSong)
}
