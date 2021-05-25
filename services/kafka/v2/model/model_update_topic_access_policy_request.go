package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTopicAccessPolicyRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *UpdateTopicAccessPolicyReq `json:"body,omitempty"`
}

func (o UpdateTopicAccessPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyRequest", string(data)}, " ")
}
