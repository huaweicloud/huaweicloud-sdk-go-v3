package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDisassociateThrottlingPolicyV2Request struct {
	ProjectId  string                      `json:"project_id"`
	InstanceId string                      `json:"instance_id"`
	Action     string                      `json:"action"`
	Body       *ThrottleBindingBatchDelete `json:"body,omitempty"`
}

func (o BatchDisassociateThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDisassociateThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateThrottlingPolicyV2Request", string(data)}, " ")
}
