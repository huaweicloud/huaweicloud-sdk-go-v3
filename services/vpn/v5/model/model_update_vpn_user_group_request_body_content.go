package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnUserGroupRequestBodyContent struct {

	// 用户组名
	Name *string `json:"name,omitempty"`

	// 用户组描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateVpnUserGroupRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserGroupRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserGroupRequestBodyContent", string(data)}, " ")
}
