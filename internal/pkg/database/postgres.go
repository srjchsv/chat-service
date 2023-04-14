package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/srjchsv/chat-service/internal/app/models"
)

func InitDB(maxRetries int, initialRetryInterval time.Duration) (*gorm.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURI := fmt.Sprintf("host=%v port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	var db *gorm.DB
	var err error
	retryInterval := initialRetryInterval

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open("postgres", dbURI)
		if err == nil {
			break
		}

		log.Printf("Failed to connect to the database (attempt %d/%d). Retrying in %v.", i+1, maxRetries, retryInterval)
		time.Sleep(retryInterval)
		retryInterval *= 2
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database after %d attempts: %w", maxRetries, err)
	}

	db.AutoMigrate(&models.Message{})

	return db, nil
}

func SaveMessage(db *gorm.DB, senderUsername, content string, createdAt time.Time) error {
	message := &models.Message{
		SenderUsername: senderUsername,
		Content:        content,
		CreatedAt:      createdAt,
	}

	if err := db.Create(message).Error; err != nil {
		return err
	}

	return nil
}

func GetMessagesFromLastDay(db *gorm.DB) ([]models.Message, error) {
	var messages []models.Message
	oneDayAgo := time.Now().Add(-24 * time.Hour)

	result := db.Where("created_at >= ?", oneDayAgo).Order("created_at DESC").Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
