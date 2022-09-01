package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTopicAccessPolicyTopicsObject struct {

	// topic名称。
	Name string `json:"name" xml:"name"`

	// 权限列表。
	Policies []UpdateTopicAccessPolicyPoliciesObject `json:"policies" xml:"policies"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 敏感字段。
	SensitiveWord *string `json:"sensitive_word,omitempty" xml:"sensitive_word"`
}

func (o UpdateTopicAccessPolicyTopicsObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyTopicsObject struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyTopicsObject", string(data)}, " ")
}
