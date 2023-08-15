package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewaySnatRuleResponse Response Object
type CreateNatGatewaySnatRuleResponse struct {
	SnatRule       *CreateNatGatewaySnatRuleResponseBody `json:"snat_rule,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o CreateNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleResponse", string(data)}, " ")
}
