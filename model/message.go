package model

import (
	"encoding/json"
	"time"
)

// Message provides a message to be distributed by RiverMQ
type Message struct {
	Timestamp time.Time       `json:"timestamp"`
	Type      string          `json:"type"`
	Body      json.RawMessage `json:"body"`
}
