package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyTopicsObject struct {

	// topic名称。
	Name string `json:"name"`

	// 权限列表。
	Policies []UpdateTopicAccessPolicyPoliciesObject `json:"policies"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 敏感字段。
	SensitiveWord *string `json:"sensitive_word,omitempty"`
}

func (o UpdateTopicAccessPolicyTopicsObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyTopicsObject struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyTopicsObject", string(data)}, " ")
}
