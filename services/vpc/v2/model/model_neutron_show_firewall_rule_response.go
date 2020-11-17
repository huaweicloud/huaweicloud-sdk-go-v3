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
type NeutronShowFirewallRuleResponse struct {
	FirewallRule *NeutronFirewallRule `json:"firewall_rule,omitempty"`
}

func (o NeutronShowFirewallRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronShowFirewallRuleResponse", string(data)}, " ")
}
