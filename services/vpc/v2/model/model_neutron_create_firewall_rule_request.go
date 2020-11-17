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
type NeutronCreateFirewallRuleRequest struct {
	Body *NeutronCreateFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronCreateFirewallRuleRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFirewallRuleRequest", string(data)}, " ")
}
