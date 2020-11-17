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

// Request Object
type NeutronUpdateFirewallGroupRequest struct {
	FirewallGroupId string                                 `json:"firewall_group_id"`
	Body            *NeutronUpdateFirewallGroupRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateFirewallGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronUpdateFirewallGroupRequest", string(data)}, " ")
}
