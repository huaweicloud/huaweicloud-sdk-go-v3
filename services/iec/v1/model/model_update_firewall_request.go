package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateFirewallRequest struct {
	// 网络ACL ID

	FirewallId string `json:"firewall_id"`

	Body *UpdateFirewallRequestBody `json:"body,omitempty"`
}

func (o UpdateFirewallRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallRequest struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRequest", string(data)}, " ")
}
