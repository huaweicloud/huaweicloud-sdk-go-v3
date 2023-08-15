package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSpecialThrottlingConfigurationV2Request Request Object
type DeleteSpecialThrottlingConfigurationV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 流控策略的编号
	ThrottleId string `json:"throttle_id"`

	// 特殊配置的编号
	StrategyId string `json:"strategy_id"`
}

func (o DeleteSpecialThrottlingConfigurationV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
