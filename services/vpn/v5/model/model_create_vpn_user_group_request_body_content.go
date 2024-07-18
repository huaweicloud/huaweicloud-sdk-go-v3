package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUserGroupRequestBodyContent struct {

	// 用户组名
	Name string `json:"name"`

	// 用户组描述
	Description *string `json:"description,omitempty"`
}

func (o CreateVpnUserGroupRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserGroupRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateVpnUserGroupRequestBodyContent", string(data)}, " ")
}
