package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSpecialThrottlingConfigurationV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 流控策略的编号

	ThrottleId string `json:"throttle_id"`
	// 特殊配置的编号

	StrategyId string `json:"strategy_id"`

	Body *ThrottleSpecialUpdate `json:"body,omitempty"`
}

func (o UpdateSpecialThrottlingConfigurationV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"UpdateSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
