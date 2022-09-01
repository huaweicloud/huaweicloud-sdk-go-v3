package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNatGatewaySnatRuleResponse struct {
	SnatRule       *NatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty" xml:"snat_rule"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleResponse", string(data)}, " ")
}
