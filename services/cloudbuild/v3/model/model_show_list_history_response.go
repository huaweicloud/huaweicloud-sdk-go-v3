package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowListHistoryResponse struct {

	// 构建历史列表
	HistoryRecords *[]HistoryRecord `json:"history_records,omitempty" xml:"history_records"`

	// 记录总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowListHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowListHistoryResponse", string(data)}, " ")
}
