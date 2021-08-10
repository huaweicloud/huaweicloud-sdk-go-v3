package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowListHistoryResponse struct {
	// 构建历史列表

	HistoryRecords *[]HistoryRecord `json:"history_records,omitempty"`
	// 记录总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowListHistoryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowListHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowListHistoryResponse", string(data)}, " ")
}
