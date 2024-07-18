package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpnUserInGroup struct {

	// 用户ID
	Id *string `json:"id,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 用户描述
	Description *string `json:"description,omitempty"`
}

func (o VpnUserInGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpnUserInGroup struct{}"
	}

	return strings.Join([]string{"VpnUserInGroup", string(data)}, " ")
}
