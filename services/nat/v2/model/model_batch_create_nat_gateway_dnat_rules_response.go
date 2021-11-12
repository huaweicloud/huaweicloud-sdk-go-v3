package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateNatGatewayDnatRulesResponse struct {
	// DNAT规则批量创建对象的响应体。

	DnatRules      *[]NatGatewayDnatRuleResponseBody `json:"dnat_rules,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchCreateNatGatewayDnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesResponse", string(data)}, " ")
}
