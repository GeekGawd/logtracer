package main

import (
	"fmt"
	"os"
	handler "github.com/GeekGawd/logtracer/internal/adapter/handler/http"
	repository "github.com/GeekGawd/logtracer/internal/adapter/repository/influxdb"

	"github.com/GeekGawd/logtracer/internal/core/service"
	"github.com/joho/godotenv"
)

func init(){
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}


func main(){

	// Create Ingest Service and Influx Adapter
	influxAdapter := repository.NewInfluxDBAdapter()
	ingestService := service.NewIngestionService(influxAdapter)
	ingestHandler := handler.NewIngestHandler(ingestService)
	defer influxAdapter.Close()
	
	router, err := handler.NewRouter(
		*ingestHandler,
	)
	if err != nil {
		fmt.Println("Could not create router: ", err)
		os.Exit(1)
	}

	listenAddr := fmt.Sprintf("127.0.0.1:%s", "8080")
	err = router.Serve(listenAddr)
	if err != nil {
		fmt.Println("Cannot start server at " + listenAddr + "Error: ", err)
		os.Exit(1)
	}
}