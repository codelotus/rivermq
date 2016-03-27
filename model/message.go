package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pborman/uuid"
)

// Message provides a message to be distributed by RiverMQ
type Message struct {
	ID        uuid.UUID       `json:"id"`
	Timestamp time.Time       `json:"timestamp"`
	Type      string          `json:"type"`
	Body      json.RawMessage `json:"body"`
}

// ValidateMessage ensures a message is valid
func ValidateMessage(msg Message) (status bool, err error) {
	if typeLen := len(msg.Type); typeLen == 0 {
		log.Printf("Error parsing supplied message Type[ %v ]\n ", msg.Type)
		return false, fmt.Errorf("Error parsing supplied message Type[ %v ]\n ", msg.Type)
	}
	return true, nil
}
