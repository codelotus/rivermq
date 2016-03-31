package model

import (
	"flag"
	"fmt"
	client "github.com/influxdata/influxdb/client/v2"
	"log"
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

// CreateRiverMQDB does that
func CreateRiverMQDB() (err error) {
	_, err = QueryDB(fmt.Sprintf("CREATE DATABASE %v", dbName))
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
