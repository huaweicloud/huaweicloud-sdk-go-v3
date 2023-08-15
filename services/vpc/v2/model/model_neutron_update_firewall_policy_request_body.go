package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateFirewallPolicyRequestBody
type NeutronUpdateFirewallPolicyRequestBody struct {
	FirewallPolicy *NeutronUpdateFirewallPolicyOption `json:"firewall_policy"`
}

func (o NeutronUpdateFirewallPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallPolicyRequestBody", string(data)}, " ")
}
