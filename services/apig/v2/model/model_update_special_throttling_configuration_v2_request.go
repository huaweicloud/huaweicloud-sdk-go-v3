package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSpecialThrottlingConfigurationV2Request struct {
	ProjectId  string                    `json:"project_id"`
	InstanceId string                    `json:"instance_id"`
	ThrottleId string                    `json:"throttle_id"`
	StrategyId string                    `json:"strategy_id"`
	Body       *ThrottleSpecialUpdateReq `json:"body,omitempty"`
}

func (o UpdateSpecialThrottlingConfigurationV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"UpdateSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
