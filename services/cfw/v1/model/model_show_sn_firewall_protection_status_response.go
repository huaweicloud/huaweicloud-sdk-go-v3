package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSnFirewallProtectionStatusResponse Response Object
type ShowSnFirewallProtectionStatusResponse struct {
	Data           *FirewallProtectionStatusVo `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowSnFirewallProtectionStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnFirewallProtectionStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSnFirewallProtectionStatusResponse", string(data)}, " ")
}
