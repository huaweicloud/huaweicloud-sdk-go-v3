package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronListFirewallRulesResponse struct {

	// firewall_rule对象列表
	FirewallRules *[]NeutronFirewallRule `json:"firewall_rules,omitempty" xml:"firewall_rules"`

	// 分页信息
	FirewallRulesLinks *[]NeutronPageLink `json:"firewall_rules_links,omitempty" xml:"firewall_rules_links"`
	HttpStatusCode     int                `json:"-"`
}

func (o NeutronListFirewallRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallRulesResponse struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallRulesResponse", string(data)}, " ")
}
