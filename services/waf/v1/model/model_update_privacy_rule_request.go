package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePrivacyRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// privacyRuleId

	RuleId string `json:"rule_id"`

	Body *UpdatePrivacyRuleRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivacyRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePrivacyRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivacyRuleRequest", string(data)}, " ")
}
