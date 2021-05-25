package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSpecialThrottlingConfigurationV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 流控策略的ID

	ThrottleId string `json:"throttle_id"`

	Body *ThrottleSpecialReq `json:"body,omitempty"`
}

func (o CreateSpecialThrottlingConfigurationV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSpecialThrottlingConfigurationV2Request struct{}"
	}

	return strings.Join([]string{"CreateSpecialThrottlingConfigurationV2Request", string(data)}, " ")
}
