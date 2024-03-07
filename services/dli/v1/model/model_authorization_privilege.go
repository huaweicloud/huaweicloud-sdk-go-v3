package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizationPrivilege 查看数据库的使用者的权限信息。
type AuthorizationPrivilege struct {

	// 授权对象，和赋权API中的“object”对应。
	Object *string `json:"object,omitempty"`

	// 判断用户是否为管理员。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 用户名称，即该用户在当前数据库上有权限。
	UserName *string `json:"user_name,omitempty"`

	// 该用户在数据库上的权限。
	Privileges *[]string `json:"privileges,omitempty"`
}

func (o AuthorizationPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizationPrivilege struct{}"
	}

	return strings.Join([]string{"AuthorizationPrivilege", string(data)}, " ")
}
