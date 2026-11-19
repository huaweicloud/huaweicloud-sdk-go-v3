package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiFirewallConfigResponse Response Object
type UpdateAiFirewallConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAiFirewallConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiFirewallConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiFirewallConfigResponse", string(data)}, " ")
}
