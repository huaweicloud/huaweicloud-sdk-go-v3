package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimelineTrafficStatisticsRequestBody struct {

	// 开始时间时间戳，毫秒时间，最多支持30天范围内的查询
	StartTime int64 `json:"start_time"`

	// 结束时间时间戳，毫秒时间
	EndTime int64 `json:"end_time"`

	// 查询时间间隔，单位为小时，范围为1-24
	Period int32 `json:"period"`

	// 资源类型，log_group / log_stream / tenant
	ResourceType string `json:"resource_type"`

	// 查询流量类型值为：write，index，storage
	SearchType string `json:"search_type"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o TimelineTrafficStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimelineTrafficStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"TimelineTrafficStatisticsRequestBody", string(data)}, " ")
}
