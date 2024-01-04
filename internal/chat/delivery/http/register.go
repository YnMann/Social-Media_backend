package http

import (
	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/api")
	{
		authEndpoints.POST("/get-contacts", h.SignUp)
	}
}
