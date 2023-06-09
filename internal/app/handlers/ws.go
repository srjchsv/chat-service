package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/srjchsv/chat-service/internal/app/models"
	"github.com/srjchsv/chat-service/internal/app/services"
	"github.com/srjchsv/chat-service/internal/pkg/database"
	"github.com/srjchsv/chat-service/internal/pkg/ws"
)

func WSHandler(db *gorm.DB, broker *kafka.Producer, counter prometheus.Counter) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		conn, err := ws.UpgradeConnection(c.Writer, c.Request)
		if err != nil {
			http.Error(c.Writer, "Failed to upgrade connection", http.StatusBadRequest)
			return
		}
		defer conn.Close()

		ws.AddClient(conn, username)

		response := models.Response{}

		for {
			// Read messages from the WebSocket and broadcast them
			var msg models.Message
			err := conn.ReadJSON(&response)
			if err != nil {
				ws.RemoveClient(conn, username)
				break
			}

			tokenString := strings.TrimPrefix(response.Token, "Bearer ")
			_, _, err = services.ValidateToken(tokenString)
			if err != nil {
				fmt.Println("Invalid token:", err)
				break
			}
			if response.Content == "" {
				continue
			}
			msg.Content = response.Content
			msg.CreatedAt = response.CreatedAt
			msg.SenderUsername = response.SenderUsername

			// Store the message in the database
			err = database.SaveMessage(db, username, msg.Content, time.Now())
			if err != nil {
				fmt.Println("Failed to store message in DB:", err)
				continue
			}

			msg.SenderUsername = username

			// Send notification to Kafka
			notification := models.Notification{
				CreatedAt:      time.Now(),
				SenderUsername: msg.SenderUsername,
				Content:        msg.Content,
			}

			notificationJSON, err := json.Marshal(notification)
			if err != nil {
				fmt.Println("Failed to marshal notification to broker:", err)
				continue
			}

			topic := "notifications"
			err = broker.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          notificationJSON,
			}, nil)

			if err != nil {
				fmt.Println("Failed to produce notification to broker:", err)
				continue
			}

			ws.BroadcastMessage(msg)
			//incretment metrics
			counter.Inc()
			time.Sleep(100 * time.Millisecond)
		}
	}
}
