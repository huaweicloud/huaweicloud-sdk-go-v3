package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateFirewallRequest struct {
	Body *CreateFirewallRequestBody `json:"body,omitempty"`
}

func (o CreateFirewallRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFirewallRequest struct{}"
	}

	return strings.Join([]string{"CreateFirewallRequest", string(data)}, " ")
}
