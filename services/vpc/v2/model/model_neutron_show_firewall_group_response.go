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
type NeutronShowFirewallGroupResponse struct {
	FirewallGroup  *NeutronFirewallGroup `json:"firewall_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronShowFirewallGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronShowFirewallGroupResponse", string(data)}, " ")
}
