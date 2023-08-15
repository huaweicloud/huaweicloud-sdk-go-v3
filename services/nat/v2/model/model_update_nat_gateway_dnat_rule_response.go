package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayDnatRuleResponse Response Object
type UpdateNatGatewayDnatRuleResponse struct {
	DnatRule       *NatGatewayDnatRuleResponseBody `json:"dnat_rule,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateNatGatewayDnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleResponse", string(data)}, " ")
}
