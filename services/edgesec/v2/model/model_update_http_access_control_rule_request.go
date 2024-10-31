package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpAccessControlRuleRequest Request Object
type UpdateHttpAccessControlRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpAccessControlRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpAccessControlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpAccessControlRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpAccessControlRuleRequest", string(data)}, " ")
}
