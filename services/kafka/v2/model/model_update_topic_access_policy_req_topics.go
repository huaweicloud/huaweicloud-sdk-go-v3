package model

import (
	"encoding/json"

	"strings"
)

// 权限实体。
type UpdateTopicAccessPolicyReqTopics struct {
	// topic名称。

	Name *string `json:"name,omitempty"`
	// 权限列表。

	Policies *[]UpdateTopicAccessPolicyReqPolicies `json:"policies,omitempty"`
}

func (o UpdateTopicAccessPolicyReqTopics) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyReqTopics struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyReqTopics", string(data)}, " ")
}
