package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRequestThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 流控策略的编号

	ThrottleId string `json:"throttle_id"`

	Body *ThrottleReq `json:"body,omitempty"`
}

func (o UpdateRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
