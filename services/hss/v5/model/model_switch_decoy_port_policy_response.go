package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchDecoyPortPolicyResponse Response Object
type SwitchDecoyPortPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchDecoyPortPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchDecoyPortPolicyResponse struct{}"
	}

	return strings.Join([]string{"SwitchDecoyPortPolicyResponse", string(data)}, " ")
}
