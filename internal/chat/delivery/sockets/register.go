package http

import (
	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(r *gin.Engine, uc chat.UseCase) {
	r.GET("/ws", func(c *gin.Context) {
		go manager.start(uc)
		wsHandler(c)
	})

	r.GET("/health", func(c *gin.Context) {
		healthHandler(c)
	})
}
