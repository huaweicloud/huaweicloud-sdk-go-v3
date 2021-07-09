package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLogsResponse struct {
	// 日志信息。

	Logs           *[]LogContents `json:"logs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
