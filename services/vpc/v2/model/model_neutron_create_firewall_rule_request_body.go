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

//
type NeutronCreateFirewallRuleRequestBody struct {
	FirewallRule *NeutronCreateFirewallRuleOption `json:"firewall_rule"`
}

func (o NeutronCreateFirewallRuleRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFirewallRuleRequestBody", string(data)}, " ")
}
