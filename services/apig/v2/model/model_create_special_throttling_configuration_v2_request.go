package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSpecialThrottlingConfigurationV2Request struct {
	InstanceId string `json:"instance_id"`

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
