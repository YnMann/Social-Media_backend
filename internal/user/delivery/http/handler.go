package http

import (
	"fmt"
	"net/http"

	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/YnMann/chat_backend/internal/models"
	"github.com/YnMann/chat_backend/internal/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUC user.UseCase
	authUC auth.UseCase
}

func NewHandler(uuc user.UseCase, auc auth.UseCase) *Handler {
	return &Handler{
		userUC: uuc,
		authUC: auc,
	}
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	// Get the user information from the context
	userToken, ok := c.Get(auth.CtxUserKey)

	fmt.Println("OK", ok)
	user, ok := userToken.(*models.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user data in context",
		})
		return
	}

	profile, err := h.userUC.GetUserProfile(c.Request.Context(), user.ID)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, profile)
}
