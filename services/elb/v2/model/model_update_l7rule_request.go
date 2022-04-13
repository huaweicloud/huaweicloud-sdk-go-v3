package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateL7ruleRequest struct {
	// 待更新的转发规则所在的转发策略id

	L7policyId string `json:"l7policy_id"`
	// 待更新的转发规则id

	L7ruleId string `json:"l7rule_id"`

	Body *UpdateL7ruleRequestBody `json:"body,omitempty"`
}

func (o UpdateL7ruleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7ruleRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleRequest", string(data)}, " ")
}
