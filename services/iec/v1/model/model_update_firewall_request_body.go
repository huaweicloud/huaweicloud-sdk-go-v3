package model

import (
	"encoding/json"

	"strings"
)

// 更新网络ACL请求体
type UpdateFirewallRequestBody struct {
	Firewall *UpdateFirewallOption `json:"firewall"`
}

func (o UpdateFirewallRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRequestBody", string(data)}, " ")
}
