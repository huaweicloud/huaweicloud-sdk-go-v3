package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteApiAclBindingV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 必须为delete

	Action string `json:"action"`

	Body *AclBindingBatchDelete `json:"body,omitempty"`
}

func (o BatchDeleteApiAclBindingV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"BatchDeleteApiAclBindingV2Request", string(data)}, " ")
}
