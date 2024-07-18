package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnUserRequestBodyContent struct {

	// 用户描述
	Description *string `json:"description,omitempty"`

	// 所属用户组ID
	UserGroupId *string `json:"user_group_id,omitempty"`
}

func (o UpdateVpnUserRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserRequestBodyContent", string(data)}, " ")
}
