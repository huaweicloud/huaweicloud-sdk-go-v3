package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteNatGatewayDnatRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteNatGatewayDnatRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNatGatewayDnatRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNatGatewayDnatRuleResponse", string(data)}, " ")
}
