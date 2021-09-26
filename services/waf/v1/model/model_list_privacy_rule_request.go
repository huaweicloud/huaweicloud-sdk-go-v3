package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPrivacyRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListPrivacyRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPrivacyRuleRequest struct{}"
	}

	return strings.Join([]string{"ListPrivacyRuleRequest", string(data)}, " ")
}
