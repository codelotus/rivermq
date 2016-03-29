package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/pborman/uuid"
)

// Message provides a message to be distributed by RiverMQ
type Message struct {
	ID        uuid.UUID       `json:"id"`
	Timestamp int64           `json:"timestamp"`
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

// SaveMessage does just that
func SaveMessage(msg Message) (resultMsg Message, err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbAddress,
		Username: dbUsername,
		Password: dbPassword,
	})
	if err != nil {
		log.Fatal(err)
		return msg, err
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "ns",
	})
	if err != nil {
		log.Fatal(err)
		return msg, err
	}

	msg.ID = uuid.NewUUID()
	msg.Timestamp = time.Now().UnixNano()

	// In influxb tags are indexed, fields are not
	tags := map[string]string{
		"Type": msg.Type,
		"ID":   msg.ID.String(),
	}
	fields := map[string]interface{}{
		"Body": msg.Body,
	}
	tm := time.Unix(0, msg.Timestamp)

	pt, err := client.NewPoint("Message", tags, fields, tm)
	if err != nil {
		log.Fatal(err)
		return msg, err
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
		return msg, err
	}

	return msg, nil
}
