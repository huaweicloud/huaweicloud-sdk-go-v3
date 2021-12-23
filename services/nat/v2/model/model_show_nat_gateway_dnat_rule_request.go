package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNatGatewayDnatRuleRequest struct {
	// DNAT规则的ID。

	DnatRuleId string `json:"dnat_rule_id"`
}

func (o ShowNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayDnatRuleRequest", string(data)}, " ")
}
