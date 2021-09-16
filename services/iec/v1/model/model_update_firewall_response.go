package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateFirewallResponse struct {
	Firewall       *UpdateFirewallResp `json:"firewall,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateFirewallResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallResponse struct{}"
	}

	return strings.Join([]string{"UpdateFirewallResponse", string(data)}, " ")
}
