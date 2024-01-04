package http

import (
	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc chat.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/api")
	{
		authEndpoints.GET("/get-contacts", h.GetContacts)
	}
}
