package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNatGatewaySnatRuleResponse Response Object
type ShowNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowNatGatewaySnatRuleResponse", string(data)}, " ")
}
