package domain

type LoggerData struct {
	Level      string `json:"level"`
	Message    string `json:"message"`
	ResourceId string `json:"resourceId"`
	Timestamp  string `json:"timestamp"`
	TraceId    string `json:"traceId"`
	SpanId     string `json:"spanId"`
	Commit     string `json:"commit"`
	Metadata   struct {
		ParentResourceId string `json:"parentResourceId"`
	} `json:"metadata"`
}

type LogQuery struct {
	Level             string `json:"level"`
	ParentResourceId  string `json:"parentResourceId"`
	ResourceId       string `json:"resourceId"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	TraceId          string `json:"traceId"`
	SpanId           string `json:"spanId"`
	Commit           string `json:"commit"`
}