package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNatGatewaySnatRulesResponse struct {

	// 查询SNAT规则列表的响应体。
	SnatRules      *[]NatGatewaySnatRuleResponseBody `json:"snat_rules,omitempty" xml:"snat_rules"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListNatGatewaySnatRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaySnatRulesResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaySnatRulesResponse", string(data)}, " ")
}
