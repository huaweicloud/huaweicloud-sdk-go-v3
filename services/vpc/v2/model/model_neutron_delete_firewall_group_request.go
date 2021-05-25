package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronDeleteFirewallGroupRequest struct {
	// 网络ACL防火墙组ID

	FirewallGroupId string `json:"firewall_group_id"`
}

func (o NeutronDeleteFirewallGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronDeleteFirewallGroupRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteFirewallGroupRequest", string(data)}, " ")
}
