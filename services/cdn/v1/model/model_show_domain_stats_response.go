package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainStatsResponse struct {
	// 数据起始时间戳，可能不返回

	StartTime *int64 `json:"start_time,omitempty"`
	// 数据结束时间戳，可能不返回

	EndTime *int64 `json:"end_time,omitempty"`
	// 查询间隔，可能不返回

	Interval *int64 `json:"interval,omitempty"`
	// 查询类型，可能不返回

	Action *string `json:"action,omitempty"`
	// 指标类型，可能不返回

	StatType *string `json:"stat_type,omitempty"`
	// 分组类型，可能不返回

	GroupBy *string `json:"group_by,omitempty"`
	// 返回数据体

	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsResponse", string(data)}, " ")
}
