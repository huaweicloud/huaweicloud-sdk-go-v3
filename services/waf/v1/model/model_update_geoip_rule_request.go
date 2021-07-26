package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateGeoipRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// geoipRuleId

	RuleId string `json:"rule_id"`

	Body *UpdateGeoipRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateGeoipRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleRequest", string(data)}, " ")
}
