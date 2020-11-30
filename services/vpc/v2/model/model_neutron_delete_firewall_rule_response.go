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
type NeutronDeleteFirewallRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteFirewallRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronDeleteFirewallRuleResponse", string(data)}, " ")
}
