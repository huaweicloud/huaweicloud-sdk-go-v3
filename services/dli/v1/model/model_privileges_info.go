package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivilegesInfo 权限信息
type PrivilegesInfo struct {

	// 判断用户是否为管理员。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 用户名称，即该用户在当前数据库上有权限。
	UserName *string `json:"user_name,omitempty"`

	// 该用户在数据库上的权限
	Privileges *[]string `json:"privileges,omitempty"`
}

func (o PrivilegesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegesInfo struct{}"
	}

	return strings.Join([]string{"PrivilegesInfo", string(data)}, " ")
}
