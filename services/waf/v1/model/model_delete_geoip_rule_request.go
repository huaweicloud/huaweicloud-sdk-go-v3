package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteGeoipRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// geoipRuleId

	RuleId string `json:"rule_id"`
}

func (o DeleteGeoipRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteGeoipRuleRequest", string(data)}, " ")
}
