package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiFirewallConfigRequest Request Object
type UpdateAiFirewallConfigRequest struct {
	Body *UpdateAiFirewallConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAiFirewallConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiFirewallConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiFirewallConfigRequest", string(data)}, " ")
}
