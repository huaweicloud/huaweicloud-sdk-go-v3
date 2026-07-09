package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbUser struct {

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 用户权限
	UserPermission *string `json:"user_permission,omitempty"`
}

func (o DbUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbUser struct{}"
	}

	return strings.Join([]string{"DbUser", string(data)}, " ")
}
