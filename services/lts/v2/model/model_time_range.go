package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeRange 此参数在请求实体中，采用json字符串格式。
type TimeRange struct {

	// 时区信息，默认为“UTC”。
	SqlTimeZone *string `json:"sql_time_zone,omitempty"`

	// 搜索起始时间（UTC时间，毫秒级）
	StartTime int64 `json:"start_time"`

	// 搜索结束时间（UTC时间，毫秒级）。
	EndTime int64 `json:"end_time"`

	// 搜索是否包含起始时间点，默认为false。
	StartTimeGt *bool `json:"start_time_gt,omitempty"`

	// 搜索是否包含结束时间点，默认为false。
	EndTimeLt *bool `json:"end_time_lt,omitempty"`
}

func (o TimeRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRange struct{}"
	}

	return strings.Join([]string{"TimeRange", string(data)}, " ")
}
