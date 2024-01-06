package http

import (
	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(r *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := r.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}
