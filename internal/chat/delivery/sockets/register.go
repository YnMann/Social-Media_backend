package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine) {
	go manager.start()

	websocketEndpoints := router.Group("/u")
	{
		websocketEndpoints.GET("/ws", func(c *gin.Context) {
			wsHandler(c.Writer, c.Request)
		})

		websocketEndpoints.GET("/health", func(c *gin.Context) {
			healthHandler(c.Writer, c.Request)
		})
	}
}
