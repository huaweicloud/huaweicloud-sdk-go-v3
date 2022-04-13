package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建DNAT规则的请求体。
type CreateNatGatewayDnatRuleOption struct {
	DnatRule *CreateNatGatewayDnatOption `json:"dnat_rule"`
}

func (o CreateNatGatewayDnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayDnatRuleOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayDnatRuleOption", string(data)}, " ")
}
