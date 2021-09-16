package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowFirewallResponse struct {
	Firewall       *Firewall `json:"firewall,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFirewallResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFirewallResponse struct{}"
	}

	return strings.Join([]string{"ShowFirewallResponse", string(data)}, " ")
}
