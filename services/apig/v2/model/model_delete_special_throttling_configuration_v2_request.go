package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSpecialThrottlingConfigurationV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	ThrottleId string `json:"throttle_id"`
	StrategyId string `json:"strategy_id"`
}

func (o DeleteSpecialThrottlingConfigurationV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
