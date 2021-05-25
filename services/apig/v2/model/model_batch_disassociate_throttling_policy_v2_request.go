package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDisassociateThrottlingPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 必须为delete

	Action string `json:"action"`

	Body *ThrottleBindingBatchDelete `json:"body,omitempty"`
}

func (o BatchDisassociateThrottlingPolicyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDisassociateThrottlingPolicyV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateThrottlingPolicyV2Request", string(data)}, " ")
}
