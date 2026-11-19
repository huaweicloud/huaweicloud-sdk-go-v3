package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiFirewallConfigRequest Request Object
type ShowAiFirewallConfigRequest struct {
}

func (o ShowAiFirewallConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiFirewallConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowAiFirewallConfigRequest", string(data)}, " ")
}
