package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateFirewallResponse struct {
	Firewall       *Firewall `json:"firewall,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateFirewallResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFirewallResponse struct{}"
	}

	return strings.Join([]string{"CreateFirewallResponse", string(data)}, " ")
}
