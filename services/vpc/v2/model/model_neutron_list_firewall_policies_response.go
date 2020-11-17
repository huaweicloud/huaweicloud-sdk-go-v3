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
type NeutronListFirewallPoliciesResponse struct {
	// firewall_policy对象列表
	FirewallPolicies *[]NeutronFirewallPolicy `json:"firewall_policies,omitempty"`
	// 分页信息
	FirewallPoliciesLinks *[]NeutronPageLink `json:"firewall_policies_links,omitempty"`
}

func (o NeutronListFirewallPoliciesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronListFirewallPoliciesResponse", string(data)}, " ")
}
