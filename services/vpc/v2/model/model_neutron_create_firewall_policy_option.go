package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NeutronCreateFirewallPolicyOption struct {

	// 审计标记。
	Audited *bool `json:"audited,omitempty" xml:"audited"`

	// 功能说明：网络ACL防火墙策略描述 取值范围：最长255个字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 策略引用的网络ACL防火墙规则链。
	FirewallRules *[]string `json:"firewall_rules,omitempty" xml:"firewall_rules"`

	// 功能说明：网络ACL防火墙策略名称 取值范围：最长255个字符
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o NeutronCreateFirewallPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallPolicyOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallPolicyOption", string(data)}, " ")
}
