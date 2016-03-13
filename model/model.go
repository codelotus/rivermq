package model

import (
	"flag"
	"fmt"
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
