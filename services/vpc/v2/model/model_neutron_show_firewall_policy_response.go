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
type NeutronShowFirewallPolicyResponse struct {
	FirewallPolicy *NeutronFirewallPolicy `json:"firewall_policy,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o NeutronShowFirewallPolicyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronShowFirewallPolicyResponse", string(data)}, " ")
}
