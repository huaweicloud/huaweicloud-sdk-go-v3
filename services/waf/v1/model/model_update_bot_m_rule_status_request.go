package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMRuleStatusRequest Request Object
type UpdateBotMRuleStatusRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	// 规则Id
	RuleId string `json:"rule_id"`

	Body *UpdateBotMRuleStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleStatusRequest", string(data)}, " ")
}
