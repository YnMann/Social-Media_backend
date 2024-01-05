package http

import (
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

func (h *Handler) GetContacts(c *gin.Context) {
	contacts, err := h.userUC.GetContacts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

func (h *Handler) GetUserProfile(c *gin.Context) {
	// Get the user information from the context
	userToken, ok := c.Get(auth.CtxUserKey)

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
