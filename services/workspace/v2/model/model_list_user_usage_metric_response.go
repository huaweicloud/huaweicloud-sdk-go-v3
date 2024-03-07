package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserUsageMetricResponse Response Object
type ListUserUsageMetricResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 用户统计指标
	Items          *[]UserMetric `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUserUsageMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserUsageMetricResponse struct{}"
	}

	return strings.Join([]string{"ListUserUsageMetricResponse", string(data)}, " ")
}
