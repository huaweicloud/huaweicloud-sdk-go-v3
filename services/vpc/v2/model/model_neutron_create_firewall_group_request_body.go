package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NeutronCreateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronCreateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronCreateFirewallGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateFirewallGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateFirewallGroupRequestBody", string(data)}, " ")
}
