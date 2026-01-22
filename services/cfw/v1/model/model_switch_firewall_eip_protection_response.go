package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchFirewallEipProtectionResponse Response Object
type SwitchFirewallEipProtectionResponse struct {
	Data *SwitchFirewallEipProtectionRespData `json:"data,omitempty"`

	// 一键逃生失败原因
	FailReason     *string `json:"fail_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchFirewallEipProtectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFirewallEipProtectionResponse struct{}"
	}

	return strings.Join([]string{"SwitchFirewallEipProtectionResponse", string(data)}, " ")
}
