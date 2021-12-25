package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowListPeriodHistoryResponse struct {
	// 记录总数

	Total *int32 `json:"total,omitempty"`
	// 构建历史列表

	HistoryRecords *[]HistoryRecord1 `json:"history_records,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowListPeriodHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListPeriodHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowListPeriodHistoryResponse", string(data)}, " ")
}
