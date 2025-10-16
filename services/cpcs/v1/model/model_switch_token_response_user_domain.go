package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseUserDomain 主账号信息
type SwitchTokenResponseUserDomain struct {

	// 主账号名称
	Name *string `json:"name,omitempty"`

	// 主账号id
	Id *string `json:"id,omitempty"`
}

func (o SwitchTokenResponseUserDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseUserDomain struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseUserDomain", string(data)}, " ")
}
