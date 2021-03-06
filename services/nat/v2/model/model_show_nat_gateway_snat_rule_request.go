package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowNatGatewaySnatRuleRequest struct {
	// SNAT规则的ID。

	SnatRuleId string `json:"snat_rule_id"`
}

func (o ShowNatGatewaySnatRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewaySnatRuleRequest", string(data)}, " ")
}
