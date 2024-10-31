package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpIgnoreRuleRequest Request Object
type UpdateHttpIgnoreRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpIgnoreRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpIgnoreRuleRequest", string(data)}, " ")
}
