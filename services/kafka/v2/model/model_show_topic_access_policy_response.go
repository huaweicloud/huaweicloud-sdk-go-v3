package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTopicAccessPolicyResponse struct {
	// topic名称。

	Name *string `json:"name,omitempty"`
	// topic类型。

	TopicType *int32 `json:"topic_type,omitempty"`
	// 权限列表。

	Policies       *[]ShowTopicAccessPolicyRespPolicies `json:"policies,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowTopicAccessPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicAccessPolicyResponse", string(data)}, " ")
}
