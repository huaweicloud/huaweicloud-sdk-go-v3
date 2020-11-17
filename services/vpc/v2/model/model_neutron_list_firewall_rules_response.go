/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronListFirewallRulesResponse struct {
	// firewall_rule对象列表
	FirewallRules *[]NeutronFirewallRule `json:"firewall_rules,omitempty"`
	// 分页信息
	FirewallRulesLinks *[]NeutronPageLink `json:"firewall_rules_links,omitempty"`
}

func (o NeutronListFirewallRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronListFirewallRulesResponse", string(data)}, " ")
}
