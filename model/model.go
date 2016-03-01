package model

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	dbName = "rivermq"
)

var (
	dbUsername string
	dbPassword string
	dbAddress  string
)

func init() {
	flag.StringVar(&dbUsername, "dbUser", "", "InfluxDB DB user")
	flag.StringVar(&dbPassword, "dbPassword", "", "InfluxDB DB password")
	flag.StringVar(&dbAddress, "dbAddress", "http://localhost:8086", "InfluxDB DB address")
}

// Subscription provides the data required to create a Subscription for clients
type Subscription struct {
	Type        string `json:"type"`
	CallbackURL string `json:"callbackUrl"`
}

// Message provides a message to be distributed by RiverMQ
type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

// Validate performs sanity checks on a Subscription instance
func (sub *Subscription) Validate() (status bool, err error) {
	if typeLen := len(sub.Type); typeLen == 0 {
		log.Printf("Error parsing supplied subscription Type[ %v ]\n ", sub.Type)
		return false, fmt.Errorf("Error parsing supplied subscription Type[ %v ]\n ", sub.Type)
	}
	_, err = url.ParseRequestURI(sub.CallbackURL)
	if err != nil {
		return false, err
	}
	return true, nil
}

// SaveSubscription does just that
func SaveSubscription(sub Subscription) (err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbAddress,
		Username: dbUsername,
		Password: dbPassword,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
		return err
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func queryDB(c client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: dbName,
	}
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
