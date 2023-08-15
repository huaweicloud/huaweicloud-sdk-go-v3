package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFirewallRuleResponse Response Object
type UpdateFirewallRuleResponse struct {
	Firewall       *UpdateFirewallRuleResp `json:"firewall,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleResponse", string(data)}, " ")
}
