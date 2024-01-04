package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(r *gin.Engine) {
	go manager.start()

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.GET("/health", func(c *gin.Context) {
		healthHandler(c.Writer, c.Request)
	})
}
