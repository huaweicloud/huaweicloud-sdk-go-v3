package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueScalingPoliciesResponse struct {

	// 策略优先级1-100，100优先级最高
	Priority *int32 `json:"priority,omitempty"`

	// 开始时间
	ImpactStartTime *string `json:"impact_start_time,omitempty"`

	// 结束时间
	ImpactStopTime *string `json:"impact_stop_time,omitempty"`

	// 最小cu数量
	MinCu *int32 `json:"min_cu,omitempty"`

	// 最大cu数量
	MaxCu *int32 `json:"max_cu,omitempty"`
}

func (o QueueScalingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueScalingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"QueueScalingPoliciesResponse", string(data)}, " ")
}
