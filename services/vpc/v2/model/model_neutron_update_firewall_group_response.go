package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateFirewallGroupResponse Response Object
type NeutronUpdateFirewallGroupResponse struct {
	FirewallGroup  *NeutronFirewallGroup `json:"firewall_group,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o NeutronUpdateFirewallGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallGroupResponse", string(data)}, " ")
}
