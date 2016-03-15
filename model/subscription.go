package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/pborman/uuid"
)

// Subscription provides the data required to create a Subscription for clients
type Subscription struct {
	ID          uuid.UUID `json:"id"`
	Timestamp   int64     `json:"timestamp"`
	Type        string    `json:"type"`
	CallbackURL string    `json:"callbackUrl"`
}

// ValidateSubscription performs sanity checks on a Subscription instance
func ValidateSubscription(sub Subscription) (status bool, err error) {
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
func SaveSubscription(sub Subscription) (resultSub Subscription, err error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     dbAddress,
		Username: dbUsername,
		Password: dbPassword,
	})
	if err != nil {
		log.Fatal(err)
		return sub, err
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbName,
		Precision: "ns",
	})
	if err != nil {
		log.Fatal(err)
		return sub, err
	}

	sub.ID = uuid.NewUUID()
	sub.Timestamp = time.Now().UnixNano()

	// In influxb tags are indexed
	tags := map[string]string{
		"Type": sub.Type,
		"ID":   sub.ID.String(),
	}
	fields := map[string]interface{}{
		"CallbackURL": sub.CallbackURL,
	}
	tm := time.Unix(0, sub.Timestamp)

	pt, err := client.NewPoint("Subscription", tags, fields, tm)
	if err != nil {
		log.Fatal(err)
		return sub, err
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Fatal(err)
		return sub, err
	}
	return sub, nil
}

// FindSubscriptionByID does what it says
func FindSubscriptionByID(id uuid.UUID) (sub Subscription, err error) {
	res, err := QueryDB("SELECT time, Type, CallbackURL, ID FROM \"Subscription\" WHERE ID = '" + id.String() + "'")
	if err != nil {
		return sub, err
	}
	if len(res) == 0 || len(res[0].Series) == 0 {
		return sub, fmt.Errorf("No Series with ID[ %s ]", id.String())
	}
	subs, err := convertResultToSubscriptionSlice(res)
	if err != nil {
		return sub, err
	}
	return subs[0], nil
}

// FindAllSubscriptionsByType does that
func FindAllSubscriptionsByType(msgType string) (subs []Subscription, err error) {
	res, err := QueryDB("select time, Type, CallbackURL, ID from \"Subscription\" where Type = '" + msgType + "'")
	if err != nil {
		return subs, err
	}
	if len(res) == 0 || len(res[0].Series) == 0 {
		return subs, fmt.Errorf("No Series with Type[ %s ]", msgType)
	}
	subs, err = convertResultToSubscriptionSlice(res)
	if err != nil {
		return subs, err
	}
	return subs, nil
}

// FindAllSubscriptions does just that
func FindAllSubscriptions() (subs []Subscription, err error) {
	res, err := QueryDB("SELECT time, Type, CallbackURL, ID FROM \"Subscription\"")
	if err != nil {
		return nil, err
	}
	//return res, err
	return convertResultToSubscriptionSlice(res)
}

// DeleteSubscriptionByID does just that
func DeleteSubscriptionByID(id uuid.UUID) (err error) {
	_, err = QueryDB("DROP SERIES FROM \"Subscription\" where ID = '" + id.String() + "'")
	return err
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
		Command:   cmd,
		Database:  dbName,
		Precision: "ns",
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

// convertResultToSubscriptionSlice does that
func convertResultToSubscriptionSlice(res []client.Result) (subs []Subscription, err error) {
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

		timeVal, _ := m["time"].(json.Number).Int64()
		sub := Subscription{
			Type:        m["Type"].(string),
			CallbackURL: m["CallbackURL"].(string),
			Timestamp:   timeVal,
			ID:          uuid.Parse(m["ID"].(string)),
		}
		result = append(result, sub)
	}

	return result, nil
}
