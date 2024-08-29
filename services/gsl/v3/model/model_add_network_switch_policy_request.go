package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNetworkSwitchPolicyRequest Request Object
type AddNetworkSwitchPolicyRequest struct {
	Body *NetworkSwitchPolicyDto `json:"body,omitempty"`
}

func (o AddNetworkSwitchPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNetworkSwitchPolicyRequest struct{}"
	}

	return strings.Join([]string{"AddNetworkSwitchPolicyRequest", string(data)}, " ")
}
