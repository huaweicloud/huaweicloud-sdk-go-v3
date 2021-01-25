package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRequestThrottlingPolicyV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	ThrottleId string `json:"throttle_id"`
}

func (o DeleteRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteRequestThrottlingPolicyV2Request", string(data)}, " ")
}
