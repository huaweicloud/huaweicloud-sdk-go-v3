package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseUsersPrivilege 查看数据库的使用者的权限信息。
type ShowDatabaseUsersPrivilege struct {

	// 判断用户是否为管理员。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 用户名称，即该用户在当前数据库上有权限。
	UserName *string `json:"user_name,omitempty"`

	// 该用户在数据库上的权限。
	Privileges *[]string `json:"privileges,omitempty"`
}

func (o ShowDatabaseUsersPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseUsersPrivilege struct{}"
	}

	return strings.Join([]string{"ShowDatabaseUsersPrivilege", string(data)}, " ")
}
