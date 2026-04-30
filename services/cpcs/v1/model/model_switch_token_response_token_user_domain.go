package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchTokenResponseTokenUserDomain 主账号信息
type SwitchTokenResponseTokenUserDomain struct {

	// 主账号名称
	Name *string `json:"name,omitempty"`

	// 主账号ID
	Id *string `json:"id,omitempty"`
}

func (o SwitchTokenResponseTokenUserDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseTokenUserDomain struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseTokenUserDomain", string(data)}, " ")
}
