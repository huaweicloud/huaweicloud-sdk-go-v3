package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronListSecurityGroupRulesResponse struct {

	// 安全组规则对象列表
	SecurityGroupRules *[]NeutronSecurityGroupRule `json:"security_group_rules,omitempty" xml:"security_group_rules"`

	// 分页信息
	SecurityGroupRulesLinks *[]NeutronPageLink `json:"security_group_rules_links,omitempty" xml:"security_group_rules_links"`
	HttpStatusCode          int                `json:"-"`
}

func (o NeutronListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupRulesResponse", string(data)}, " ")
}
