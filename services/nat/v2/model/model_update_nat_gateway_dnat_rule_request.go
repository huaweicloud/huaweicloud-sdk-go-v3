package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNatGatewayDnatRuleRequest struct {
	// DNAT规则的ID。

	DnatRuleId string `json:"dnat_rule_id"`

	Body *UpdateNatGatewayDnatRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateNatGatewayDnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequest", string(data)}, " ")
}
