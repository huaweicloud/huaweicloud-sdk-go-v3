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
type NeutronUpdateFirewallGroupResponse struct {
	FirewallGroup *NeutronFirewallGroup `json:"firewall_group,omitempty"`
}

func (o NeutronUpdateFirewallGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFirewallGroupResponse", string(data)}, " ")
}
