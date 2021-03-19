package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRequestThrottlingPolicyV2Request struct {
	InstanceId string `json:"instance_id"`

	Body *ThrottleReq `json:"body,omitempty"`
}

func (o CreateRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"CreateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
