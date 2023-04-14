package routes

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jinzhu/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/chat-service/internal/app/handlers"
)

func SetupRouter(db *gorm.DB, producer *kafka.Producer) *gin.Engine {
	router := gin.Default()

	// Allow all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	// Use the CORS middleware
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	//v1.Use(handlers.AuthMiddleware())
	{
		// v1.POST("/messages", handlers.SendMessage(db))
		v1.GET("/ws/:username", handlers.WSHandler(db, producer))
		v1.GET("/chat_history", handlers.GetChatHistoryHandler(db))
	}

	return router
}
