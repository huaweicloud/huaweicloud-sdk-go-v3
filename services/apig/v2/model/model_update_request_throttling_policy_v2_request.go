package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRequestThrottlingPolicyV2Request struct {
	ProjectId  string       `json:"project_id"`
	InstanceId string       `json:"instance_id"`
	ThrottleId string       `json:"throttle_id"`
	Body       *ThrottleReq `json:"body,omitempty"`
}

func (o UpdateRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
