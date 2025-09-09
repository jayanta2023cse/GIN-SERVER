// Package websocket handles application configuration, including loading environment variables.
package websocket

import (
	// Import your helpers package (adjust path)
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// WebSocketHandler handles WebSocket connections without authentication
func WebSocketHandler(c *gin.Context) {
	// Get the topic from URL param (e.g., /ws/:topic)
	topic := c.Param("topic")
	if topic == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing topic",
			"status":  false,
			"message": "Topic parameter is required",
		})
		return
	}

	// Temporary clientID (e.g., from query param or generate UUID; replace with auth later)
	clientID := c.Query("clientID") // Example: pass ?clientID=123 in URL for testing
	if clientID == "" {
		clientID = "anonymous-" + uuid.NewString() // Use github.com/google/uuid for random ID
	}

	// Upgrade to WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // Allow all; restrict in prod
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to upgrade connection",
			"status":  false,
			"message": "WebSocket upgrade failed",
		})
		return
	}

	// Add connection to topic
	AddConnection(topic, clientID, conn)

	// Handle disconnection
	defer func() {
		RemoveConnection(topic, clientID)
		conn.Close()
		log.Printf("Client %s disconnected from topic %s", clientID, topic)
	}()

	log.Printf("Client %s connected to topic %s", clientID, topic)

	// Main loop: Listen for incoming messages from client (vice versa)
	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Connection error with client %s: %v", clientID, err)
			break
		}
		// Process incoming message (e.g., broadcast back or send to Kafka)
		log.Printf("Received message from client %s in topic %s: %s", clientID, topic, messageBytes)

		// Example: Broadcast the received message to all in the topic
		var msg map[string]interface{}
		json.Unmarshal(messageBytes, &msg) // Assume JSON
		EmitMessageToTopic(topic, msg)
	}
}
