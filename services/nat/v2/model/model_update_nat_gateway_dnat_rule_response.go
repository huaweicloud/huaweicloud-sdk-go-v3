package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNatGatewayDnatRuleResponse struct {
	DnatRule       *NatGatewayDnatRuleResponseBody `json:"dnat_rule,omitempty" xml:"dnat_rule"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateNatGatewayDnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleResponse", string(data)}, " ")
}
