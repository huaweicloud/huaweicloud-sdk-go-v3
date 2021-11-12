package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowListPeriodHistoryResponse struct {
	// 构建历史列表

	HistoryRecords *[]HistoryRecord `json:"history_records,omitempty"`
	// 记录总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowListPeriodHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListPeriodHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowListPeriodHistoryResponse", string(data)}, " ")
}
