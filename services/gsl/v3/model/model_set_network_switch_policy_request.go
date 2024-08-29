package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetNetworkSwitchPolicyRequest Request Object
type SetNetworkSwitchPolicyRequest struct {

	// SIM卡标识
	SimCardId int64 `json:"sim_card_id"`

	Body *NetworkSwitchPolicyReq `json:"body,omitempty"`
}

func (o SetNetworkSwitchPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetNetworkSwitchPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetNetworkSwitchPolicyRequest", string(data)}, " ")
}
