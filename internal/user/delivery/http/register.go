package http

import (
	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/YnMann/chat_backend/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uuc user.UseCase, auc auth.UseCase) {
	h := NewHandler(uuc, auc)

	authEndpoints := router.Group("/user")
	{
		authEndpoints.GET("/get-profile", h.GetUserProfile)
		authEndpoints.GET("/get-contacts", h.GetContacts)
	}
}
