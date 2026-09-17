package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogStatisticsNewRequestBody 获取慢日志统计请求体
type ShowSlowLogStatisticsNewRequestBody struct {

	// 节点ID列表
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 统计字段，取值范围：nodeId、sqlType、dbName、collection、user、client
	StatisticsField string `json:"statistics_field"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序（true：正序，false：逆序）
	SortAsc *bool `json:"sort_asc,omitempty"`
}

func (o ShowSlowLogStatisticsNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsNewRequestBody struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsNewRequestBody", string(data)}, " ")
}
