package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableUsersRespPrivilege 查看表的使用者响应的权限信息。
type ShowTableUsersRespPrivilege struct {

	// 判断是否为管理用户。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 该用户有权限的对象： “databases.数据库名.tables.表名”，用户在当前表上的权限。 “databases.数据库名.tables.表名.columns.列名”，用户在列上的权限。
	Object *string `json:"object,omitempty"`

	// 该用户在相应object上的权限。
	Privileges *[]string `json:"privileges,omitempty"`

	// 拥有该权限的用户名。
	UserName *string `json:"user_name,omitempty"`
}

func (o ShowTableUsersRespPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableUsersRespPrivilege struct{}"
	}

	return strings.Join([]string{"ShowTableUsersRespPrivilege", string(data)}, " ")
}
