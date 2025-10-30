package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFirewallStatusResponse Response Object
type SwitchFirewallStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchFirewallStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFirewallStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchFirewallStatusResponse", string(data)}, " ")
}
