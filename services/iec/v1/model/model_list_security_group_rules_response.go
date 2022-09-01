package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecurityGroupRulesResponse struct {

	// 安全组规则列表对象。
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules,omitempty" xml:"security_group_rules"`

	// 安全组规则数目。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityGroupRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesResponse", string(data)}, " ")
}
