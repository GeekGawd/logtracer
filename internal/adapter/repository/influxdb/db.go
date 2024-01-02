package repository

import (
	"fmt"
	"os"
	"context"
	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/GeekGawd/logtracer/internal/core/domain"
)


type InfluxDBAdapter struct {
	client *influxdb3.Client
}

func NewInfluxDBAdapter() *InfluxDBAdapter {
	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")
	// client := influxdb2.NewClient(url, token)
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:  url,
		Token: token,
		Database: database,
		Organization: org,
	})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: " + err.Error())
		os.Exit(1)
	}

	return &InfluxDBAdapter{client: client}
}


func (a *InfluxDBAdapter) Insert(data domain.LoggerData) error {

	line := fmt.Sprintf(`logs,level="%s",resourceId="%s",traceId="%s",spanId="%s",commit="%s",parentResourceId="%s" message="%s"`, data.Level, data.ResourceId, data.TraceId, data.SpanId, data.Commit, data.Metadata.ParentResourceId, data.Message)

	err := a.client.Write(context.Background(), []byte(line))
	if err != nil {
		return err
	}
	return nil
}

func (a *InfluxDBAdapter) BulkInsert(data []domain.LoggerData) error {
	// bucket := os.Getenv("INFLUXDB_BUCKET")
	// org := os.Getenv("INFLUXDB_ORG")

	// writeAPI := a.client.WriteAPI(org, bucket)
	// defer writeAPI.Close()

	// for _, d := range data {
	// 	line := fmt.Sprintf(`logs,level="%s",resourceId="%s",traceId="%s",spanId="%s",commit="%s",parentResourceId="%s" message="%s"`, d.Level, d.ResourceId, d.TraceId, d.SpanId, d.Commit, d.Metadata.ParentResourceId, d.Message)
	// 	writeAPI.WriteRecord(line)
	// }

	// // Flush any remaining data
	// writeAPI.Flush()

	return nil
}


func (a *InfluxDBAdapter) Query(query domain.LogQuery) ([]domain.LoggerData, error) {
	// error not implemented
	return nil, nil
}

func (a *InfluxDBAdapter) Close(){
	err := a.client.Close()
	if err != nil {
		fmt.Println("Error closing InfluxDB Client: " + err.Error())
		os.Exit(1)
	}
}