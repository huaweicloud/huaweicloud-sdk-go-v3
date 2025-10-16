package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateIpReputationRuleRequestBody struct {

	// name
	Name *string `json:"name,omitempty"`

	// type
	Type *string `json:"type,omitempty"`

	// tags
	Tags *[]string `json:"tags,omitempty"`

	Action *BatchCreateIpReputationRuleRequestBodyAction `json:"action,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// 策略id
	PolicyIds *[]string `json:"policy_ids,omitempty"`
}

func (o BatchCreateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateIpReputationRuleRequestBody", string(data)}, " ")
}
