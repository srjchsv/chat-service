package routes

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/jinzhu/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/srjchsv/chat-service/internal/app/handlers"
	"github.com/srjchsv/chat-service/internal/pkg/appmetrics"
)

func SetupRouter(db *gorm.DB, producer *kafka.Producer) *gin.Engine {
	router := gin.Default()

	// Allow all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	// Use the CORS middleware
	router.Use(cors.New(config))
	// Init metrics
	metricsPrometheus := appmetrics.InitPrometheus(router)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ws/:username", handlers.WSHandler(db, producer, metricsPrometheus.WsChatMessagesSent))
		v1.GET("/chat_history", handlers.AuthMiddleware(), handlers.GetChatHistoryHandler(db))
	}

	return router
}
