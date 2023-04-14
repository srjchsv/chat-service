package ws

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/srjchsv/chat-service/internal/app/models"
)

var clients = make(map[*websocket.Conn]string)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func AddClient(conn *websocket.Conn, username string) {
	clients[conn] = username
	fmt.Printf("New client added: %s\n", username)
}

func RemoveClient(conn *websocket.Conn, username string) {
	delete(clients, conn)
	fmt.Printf("Client removed: %s\n", username)
}

func SendMessage(senderID, receiverID uint, content string) error {
	for conn, username := range clients {
		if username == strconv.Itoa(int(receiverID)) {
			err := conn.WriteJSON(map[string]interface{}{
				"sender_id": senderID,
				"content":   content,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func BroadcastMessage(msg models.Message) {
	for conn := range clients {
		err := conn.WriteJSON(msg)
		if err != nil {
			conn.Close()
			delete(clients, conn)
		}
	}
}
