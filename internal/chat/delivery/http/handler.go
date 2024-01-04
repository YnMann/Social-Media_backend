package http

import (
	"net/http"

	"github.com/YnMann/chat_backend/internal/auth"
	"github.com/YnMann/chat_backend/internal/chat"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase chat.UseCase
}

func NewHandler(uc auth.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) GetContacts(c *gin.Context) {
	c, err := h.useCase.GetContacts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}
}
