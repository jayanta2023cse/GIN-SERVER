// Package websocket handles application configuration, including loading environment variables.
package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// ConnectionsByTopic organizes WebSocket connections by topic
var (
	ConnectionsByTopic = make(map[string]map[string]*websocket.Conn)
	ConnMutex          = sync.Mutex{}
)

// AddConnection adds a new connection to a specific topic
func AddConnection(topic, clientID string, conn *websocket.Conn) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	if ConnectionsByTopic[topic] == nil {
		ConnectionsByTopic[topic] = make(map[string]*websocket.Conn)
	}
	ConnectionsByTopic[topic][clientID] = conn

	log.Printf("Connection added: clientID=%s, topic=%s", clientID, topic)
}

// RemoveConnection removes a connection from a specific topic
func RemoveConnection(topic, clientID string) {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	if ConnectionsByTopic[topic] != nil {
		delete(ConnectionsByTopic[topic], clientID)
		log.Printf("Connection removed: clientID=%s, topic=%s", clientID, topic)

		// If the topic is now empty, clean up the map
		if len(ConnectionsByTopic[topic]) == 0 {
			delete(ConnectionsByTopic, topic)
			log.Printf("Topic deleted: topic=%s (no more connections)", topic)
		}
	}
}

// EmitMessageToTopic sends a message to all connections in a specific topic
func EmitMessageToTopic(topic string, message interface{}) error {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	connections, ok := ConnectionsByTopic[topic]
	if !ok {
		log.Printf("Emit failed: topic=%s does not exist", topic)
		return fmt.Errorf("internal error:topic not found")
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	for clientID, conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			log.Printf("Failed to send message to clientID=%s in topic=%s: %v", clientID, topic, err)
			conn.Close()
			delete(connections, clientID)
		} else {
			log.Printf("Message sent to clientID=%s in topic=%s", clientID, topic)
		}
	}

	// Clean up the topic if it's now empty
	if len(connections) == 0 {
		delete(ConnectionsByTopic, topic)
		log.Printf("Topic deleted: topic=%s (no more connections)", topic)
	}

	return nil
}

// EmitMessageToClient sends a message to a specific client in a topic
func EmitMessageToClient(topic string, clientID string, message interface{}) error {
	ConnMutex.Lock()
	defer ConnMutex.Unlock()

	// Check if the topic exists
	connections, ok := ConnectionsByTopic[topic]
	if !ok {
		return fmt.Errorf("topic %s does not exist", topic)
	}

	// Check if the client exists in the topic
	conn, ok := connections[clientID]
	if !ok {
		return fmt.Errorf("client %s does not exist in topic %s", clientID, topic)
	}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// Send the message to the specific client
	if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
		log.Printf("Failed to send message to clientID=%s in topic=%s: %v", clientID, topic, err)
		conn.Close()
		delete(connections, clientID)
		return fmt.Errorf("failed to send message to clientID=%s: %v", clientID, err)
	}

	log.Printf("Message sent to clientID=%s in topic=%s", clientID, topic)
	return nil
}
