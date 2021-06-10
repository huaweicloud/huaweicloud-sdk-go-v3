package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLogsResponse struct {
	// 总数。

	Total *int32 `json:"total,omitempty"`
	// 日志列表数据

	Logs           *[]LogObject `json:"logs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowLogsResponse struct{}"
	}

	return strings.Join([]string{"ShowLogsResponse", string(data)}, " ")
}
