package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpCcRuleRequest Request Object
type UpdateHttpCcRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// cc规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpCcRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpCcRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpCcRuleRequest", string(data)}, " ")
}
