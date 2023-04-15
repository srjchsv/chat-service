package models

import "time"

type Response struct {
	Token string `json:"token"`
	Message
}

type Message struct {
	CreatedAt      time.Time `json:"created_at"`
	SenderID       uint      `json:"sender_id"`
	SenderUsername string    `json:"sender_username"`
	Content        string    `json:"content"`
}

type Notification struct {
	CreatedAt      time.Time `json:"created_at"`
	SenderUsername string    `json:"sender_username"`
	Content        string    `json:"content"`
}
