package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateNatGatewayDnatRuleRequest struct {
	// DNAT规则的ID。

	DnatRuleId string `json:"dnat_rule_id"`

	Body *UpdateNatGatewayDnatRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayDnatRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequest", string(data)}, " ")
}
