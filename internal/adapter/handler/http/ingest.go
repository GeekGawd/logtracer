package handler

import (
	"github.com/GeekGawd/logtracer/internal/core/domain"
	"github.com/GeekGawd/logtracer/internal/core/port"
	"github.com/gin-gonic/gin"
)

type IngestHandler struct {
	service port.IngestionPort
}

func NewIngestHandler(service port.IngestionPort) *IngestHandler {
	return &IngestHandler{service: service}
}

func (ig *IngestHandler) HelloWorld(ctx *gin.Context) {
	handleSuccess(ctx, "Hello World")
}

func (ig *IngestHandler) IngestData(ctx *gin.Context) {
	var data domain.LoggerData

	err := ctx.ShouldBindJSON(&data); if err != nil{
		validationError(ctx, err)
		return
	}
	
	err = ig.service.Insert(data)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "Data Ingested Successfully")
}

func (ig *IngestHandler) QueryData(ctx *gin.Context) {
	var query domain.LogQuery

	err := ctx.ShouldBindJSON(&query); if err != nil{
		validationError(ctx, err)
		return
	}
	
	data, err := ig.service.Query(query)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, data)
}

func (ig *IngestHandler) BulkIngestData(ctx *gin.Context) {
	var data []domain.LoggerData

	err := ctx.ShouldBindJSON(&data); if err != nil{
		validationError(ctx, err)
		return
	}
	
	err = ig.service.BulkInsert(data)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, "Data Ingested Successfully")
}

