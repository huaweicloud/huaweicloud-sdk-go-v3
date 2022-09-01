package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty" xml:"snat_rule"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowNatGatewaySnatRuleResponse", string(data)}, " ")
}
