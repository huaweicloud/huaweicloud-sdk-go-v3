package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHistoryTasksResponse struct {
	// 总数。

	Total *int32 `json:"total,omitempty"`
	// 日志列表数据

	Tasks          *[]TasksObject `json:"tasks,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowHistoryTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHistoryTasksResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTasksResponse", string(data)}, " ")
}
