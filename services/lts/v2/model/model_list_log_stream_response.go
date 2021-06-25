package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListLogStreamResponse struct {
	// 日志组数组。

	LogStreams     *[]LogStream `json:"log_streams,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListLogStreamResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLogStreamResponse struct{}"
	}

	return strings.Join([]string{"ListLogStreamResponse", string(data)}, " ")
}
