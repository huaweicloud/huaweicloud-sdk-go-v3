package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAppWhitelistPolicyLearnStatusResponse Response Object
type SwitchAppWhitelistPolicyLearnStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchAppWhitelistPolicyLearnStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAppWhitelistPolicyLearnStatusResponse struct{}"
	}

	return strings.Join([]string{"SwitchAppWhitelistPolicyLearnStatusResponse", string(data)}, " ")
}
