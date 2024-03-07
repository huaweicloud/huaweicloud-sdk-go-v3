package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopUsageMetricResponse Response Object
type ListDesktopUsageMetricResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 桌面统计指标
	Items          *[]DesktopMetric `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDesktopUsageMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopUsageMetricResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopUsageMetricResponse", string(data)}, " ")
}
