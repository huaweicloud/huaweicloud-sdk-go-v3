package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新DNAT规则的响应体。
type UpdateNatGatewayDnatRuleRequestBody struct {
	DnatRule *UpdateNatGatewayDnatRuleOption `json:"dnat_rule,omitempty"`
}

func (o UpdateNatGatewayDnatRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleRequestBody", string(data)}, " ")
}
