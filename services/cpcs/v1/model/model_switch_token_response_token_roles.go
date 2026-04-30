package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchTokenResponseTokenRoles struct {

	// 角色名称
	Name *string `json:"name,omitempty"`

	// 角色ID
	Id *int32 `json:"id,omitempty"`
}

func (o SwitchTokenResponseTokenRoles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchTokenResponseTokenRoles struct{}"
	}

	return strings.Join([]string{"SwitchTokenResponseTokenRoles", string(data)}, " ")
}
