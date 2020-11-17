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
type NeutronUpdateFirewallPolicyResponse struct {
	FirewallPolicy *NeutronFirewallPolicy `json:"firewall_policy,omitempty"`
}

func (o NeutronUpdateFirewallPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFirewallPolicyResponse", string(data)}, " ")
}
