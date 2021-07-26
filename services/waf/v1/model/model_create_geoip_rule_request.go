package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateGeoipRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *CreateGeoIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateGeoipRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGeoipRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateGeoipRuleRequest", string(data)}, " ")
}
