package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNetworkSwitchPolicyResponse Response Object
type AddNetworkSwitchPolicyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddNetworkSwitchPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNetworkSwitchPolicyResponse struct{}"
	}

	return strings.Join([]string{"AddNetworkSwitchPolicyResponse", string(data)}, " ")
}
