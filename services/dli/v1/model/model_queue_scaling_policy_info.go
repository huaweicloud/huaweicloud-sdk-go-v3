package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueueScalingPolicyInfo 队列优先级策略信息
type QueueScalingPolicyInfo struct {

	// 优先级
	Priority int32 `json:"priority"`

	// 开始时间
	ImpactStartTime string `json:"impact_start_time"`

	// 结束时间
	ImpactStopTime string `json:"impact_stop_time"`

	// 最小cu数量
	MinCu int32 `json:"min_cu"`

	// 最大cu数量
	MaxCu int32 `json:"max_cu"`
}

func (o QueueScalingPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueScalingPolicyInfo struct{}"
	}

	return strings.Join([]string{"QueueScalingPolicyInfo", string(data)}, " ")
}
