package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/srjchsv/chat-service/internal/pkg/database"
)

func GetChatHistoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		messages, err := database.GetMessagesFromLastDay(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve chat history"})
			return
		}

		c.JSON(http.StatusOK, messages)
	}
}
