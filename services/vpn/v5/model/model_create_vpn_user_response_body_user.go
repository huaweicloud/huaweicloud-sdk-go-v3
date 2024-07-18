package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUserResponseBodyUser struct {

	// VPN用户ID
	Id *string `json:"id,omitempty"`
}

func (o CreateVpnUserResponseBodyUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserResponseBodyUser struct{}"
	}

	return strings.Join([]string{"CreateVpnUserResponseBodyUser", string(data)}, " ")
}
