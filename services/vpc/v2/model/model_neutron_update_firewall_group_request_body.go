package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateFirewallGroupRequestBody
type NeutronUpdateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronUpdateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronUpdateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupRequestBody", string(data)}, " ")
}
