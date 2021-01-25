package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfRequestThrottlingPolicyV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	ThrottleId string `json:"throttle_id"`
}

func (o ShowDetailsOfRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfRequestThrottlingPolicyV2Request", string(data)}, " ")
}
