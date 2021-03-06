package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSpecialThrottlingConfigurationV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 流控策略的ID

	ThrottleId string `json:"throttle_id"`
	// 特殊配置的编号

	StrategyId string `json:"strategy_id"`
}

func (o DeleteSpecialThrottlingConfigurationV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
