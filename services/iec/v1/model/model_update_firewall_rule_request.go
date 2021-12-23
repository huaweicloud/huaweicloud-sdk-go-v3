package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFirewallRuleRequest struct {
	// 网络ACL ID

	FirewallId string `json:"firewall_id"`

	Body *UpdateFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleRequest", string(data)}, " ")
}
