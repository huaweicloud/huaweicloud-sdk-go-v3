package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNatGatewayDnatRulesResponse struct {
	// 查询DNAT规则列表的响应体。

	DnatRules      *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListNatGatewayDnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayDnatRulesResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewayDnatRulesResponse", string(data)}, " ")
}
