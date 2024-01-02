# Log Aggregation System with Golang, Grafana and InfluxDB using Hexagonal Architecture

So, this is a log aggregation system which takes in the json data using an endpoint.

You can do a post request to the following endpoint `http://localhost:3000/` with the following sample data: 

```
{
	"level": "error",
	"message": "Failed to connect to DB",
    "resourceId": "server-1234",
	"timestamp": "2023-09-15T08:00:00Z",
	"traceId": "abc-xyz-123",
    "spanId": "span-456",
    "commit": "5e5342f",
    "metadata": {
        "parentResourceId": "server-0987"
    }
}
```