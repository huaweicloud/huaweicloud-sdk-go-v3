package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateFirewallRuleResponse struct {
	Firewall       *UpdateFirewallRuleResp `json:"firewall,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateFirewallRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleResponse", string(data)}, " ")
}
