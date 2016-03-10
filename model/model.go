package model

import (
	"encoding/json"
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
	Timestamp   string `json:"timestamp"`
	Type        string `json:"type"`
	CallbackURL string `json:"callbackUrl"`
}

// Message provides a message to be distributed by RiverMQ
type Message struct {
	Timestamp time.Time       `json:"timestamp"`
	Type      string          `json:"type"`
	Body      json.RawMessage `json:"body"`
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

// CreateRiverMQDB does that
func CreateRiverMQDB() (err error) {
	_, err = QueryDB(fmt.Sprintf("CREATE DATABASE %v", dbName))
	return err
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

	tags := map[string]string{"Type": sub.Type}
	fields := map[string]interface{}{
		"CallbackURL": sub.CallbackURL,
	}

	pt, err := client.NewPoint("Subscription", tags, fields, time.Now())
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

// GetAllSubscriptions does just that
//func GetAllSubscriptions() (res []client.Result, err error) {
func GetAllSubscriptions() (subs []Subscription, err error) {
	res, err := QueryDB("select time, Type, CallbackURL from \"Subscription\"")
	if err != nil {
		return nil, err
	}
	//return res, err
	return ConvertResultToSubscriptionSlice(res)
}

// QueryDB does that
func QueryDB(cmd string) (res []client.Result, err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbAddress,
		Username: dbUsername,
		Password: dbPassword,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer c.Close()
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

// ConvertResultToSubscriptionSlice does that
func ConvertResultToSubscriptionSlice(res []client.Result) (subs []Subscription, err error) {
	series := res[0].Series[0]
	var result []Subscription
	columns := series.Columns

	// TODO: Figure out how to do this without a temporary map
	for i := range series.Values {
		values := series.Values[i]
		m := make(map[string]interface{})
		for x := range values {
			m[columns[x]] = values[x]
		}
		sub := Subscription{
			Type:        m["Type"].(string),
			CallbackURL: m["CallbackURL"].(string),
			Timestamp:   m["time"].(string),
		}
		result = append(result, sub)
	}

	return result, nil
}
