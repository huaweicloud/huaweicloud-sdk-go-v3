package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseUser 用户信息
type SwitchTokenResponseUser struct {
	Domain *SwitchTokenResponseUserDomain `json:"domain,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 用户id
	Id *string `json:"id,omitempty"`
}

func (o SwitchTokenResponseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseUser struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseUser", string(data)}, " ")
}
