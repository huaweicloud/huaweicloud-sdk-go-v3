package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePrivacyRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *CreatePrivacyRuleRequestBody `json:"body,omitempty"`
}

func (o CreatePrivacyRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePrivacyRuleRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivacyRuleRequest", string(data)}, " ")
}
