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
type NeutronUpdateFirewallRuleResponse struct {
	FirewallRule   *NeutronFirewallRule `json:"firewall_rule,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o NeutronUpdateFirewallRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFirewallRuleResponse", string(data)}, " ")
}
