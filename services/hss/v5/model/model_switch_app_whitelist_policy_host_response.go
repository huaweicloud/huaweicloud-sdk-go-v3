package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAppWhitelistPolicyHostResponse Response Object
type SwitchAppWhitelistPolicyHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchAppWhitelistPolicyHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAppWhitelistPolicyHostResponse struct{}"
	}

	return strings.Join([]string{"SwitchAppWhitelistPolicyHostResponse", string(data)}, " ")
}
