package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNatGatewaySnatRuleResponse Response Object
type DeleteNatGatewaySnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewaySnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewaySnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewaySnatRuleResponse", string(data)}, " ")
}
