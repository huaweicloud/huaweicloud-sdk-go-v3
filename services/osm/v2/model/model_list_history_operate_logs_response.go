package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListHistoryOperateLogsResponse struct {
	// 总记录数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 操作列表

	OpsList        *[]OperateLog `json:"ops_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListHistoryOperateLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHistoryOperateLogsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryOperateLogsResponse", string(data)}, " ")
}
