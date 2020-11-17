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
type NeutronCreateFirewallGroupResponse struct {
	FirewallGroup *NeutronFirewallGroup `json:"firewall_group,omitempty"`
}

func (o NeutronCreateFirewallGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFirewallGroupResponse", string(data)}, " ")
}
