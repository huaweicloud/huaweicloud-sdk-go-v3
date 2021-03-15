package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateRequestThrottlingPolicyV2Request struct {
	InstanceId        string `json:"instance_id"`
	ThrottleBindingId string `json:"throttle_binding_id"`
}

func (o DisassociateRequestThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateRequestThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateRequestThrottlingPolicyV2Request", string(data)}, " ")
}
