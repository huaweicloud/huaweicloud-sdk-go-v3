package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateL7RuleRequest struct {
	// 策略ID。

	L7policyId string `json:"l7policy_id"`
	// 规则ID。

	L7ruleId string `json:"l7rule_id"`

	Body *UpdateL7RuleRequestBody `json:"body,omitempty"`
}

func (o UpdateL7RuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequest", string(data)}, " ")
}
