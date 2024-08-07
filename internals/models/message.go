package models

import "github.com/gorilla/websocket"

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type RoomMessage struct {
	RoomID  string          `json:"room_id"`
	Message Message         `json:"message"`
	Sender  *websocket.Conn `json:"sender"`
}
