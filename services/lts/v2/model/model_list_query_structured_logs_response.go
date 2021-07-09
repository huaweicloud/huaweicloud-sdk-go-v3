package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListQueryStructuredLogsResponse struct {
	// 日志信息。

	StructLogs     *[]StructLogContents `json:"struct_logs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListQueryStructuredLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQueryStructuredLogsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryStructuredLogsResponse", string(data)}, " ")
}
