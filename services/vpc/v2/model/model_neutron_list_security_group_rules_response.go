package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronListSecurityGroupRulesResponse Response Object
type NeutronListSecurityGroupRulesResponse struct {

	// 安全组规则对象列表
	SecurityGroupRules *[]NeutronSecurityGroupRule `json:"security_group_rules,omitempty"`

	// 分页信息
	SecurityGroupRulesLinks *[]NeutronPageLink `json:"security_group_rules_links,omitempty"`
	HttpStatusCode          int                `json:"-"`
}

func (o NeutronListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupRulesResponse", string(data)}, " ")
}
