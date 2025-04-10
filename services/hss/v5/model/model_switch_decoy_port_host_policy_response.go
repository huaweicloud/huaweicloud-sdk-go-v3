package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchDecoyPortHostPolicyResponse Response Object
type SwitchDecoyPortHostPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchDecoyPortHostPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchDecoyPortHostPolicyResponse struct{}"
	}

	return strings.Join([]string{"SwitchDecoyPortHostPolicyResponse", string(data)}, " ")
}
