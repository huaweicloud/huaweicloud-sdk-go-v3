package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteNatGatewaySnatRuleRequest struct {
	// 公网NAT网关实例的ID。

	NatGatewayId string `json:"nat_gateway_id"`
	// SNAT规则的ID。

	SnatRuleId string `json:"snat_rule_id"`
}

func (o DeleteNatGatewaySnatRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewaySnatRuleRequest", string(data)}, " ")
}
