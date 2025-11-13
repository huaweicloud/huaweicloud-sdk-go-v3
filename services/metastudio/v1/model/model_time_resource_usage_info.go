package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeResourceUsageInfo struct {

	// 查询时间段开始时间,格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	Time *string `json:"time,omitempty"`

	// 资源用量列表
	ResourcesUsage *[]ResourceUsageInfo `json:"resources_usage,omitempty"`
}

func (o TimeResourceUsageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeResourceUsageInfo struct{}"
	}

	return strings.Join([]string{"TimeResourceUsageInfo", string(data)}, " ")
}
