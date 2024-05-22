package database

import (
	"log"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func ConnectToDB() (*influxdb3.Client, error) {
	url := "https://us-east-1-1.aws.cloud2.influxdata.com"
	token := "MDjBexPHuWxo15wb9Y_KFTp1LPPCG3kRin7mktRMzGRDbPb0C6m22NX40kVUm7tZw2vGwa3zUQ0MjWdUm6d5OA=="
	database := "smartcampusmaua"

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return &influxdb3.Client{}, err
	}

	return influxdb3Client, nil
}
