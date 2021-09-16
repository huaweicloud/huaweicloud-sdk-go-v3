package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFirewallsResponse struct {
	// 网络ACL数量。

	Count *int32 `json:"count,omitempty"`
	// 网络ACL列表。

	Firewalls      *[]Firewall `json:"firewalls,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListFirewallsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFirewallsResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallsResponse", string(data)}, " ")
}
