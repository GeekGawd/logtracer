package handler

import (
	"github.com/gin-gonic/gin"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

func NewRouter(
	ingestHandler IngestHandler,
) (*Router, error) {

	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", ingestHandler.HelloWorld)
		v1.POST("/ingest", ingestHandler.IngestData)
		v1.POST("/bulk-ingest", ingestHandler.BulkIngestData)
	}

	return &Router{router}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
	
