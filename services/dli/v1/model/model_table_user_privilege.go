package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableUserPrivilege 查看表的使用者响应的权限信息。
type TableUserPrivilege struct {

	// 判断是否为管理用户。
	IsAdmin *bool `json:"is_admin,omitempty"`

	// 该用户有权限的对象： “databases.数据库名.tables.表名”，用户在当前表上的权限。 “databases.数据库名.tables.表名.columns.列名”，用户在列上的权限。
	Object *string `json:"object,omitempty"`

	// 该用户在相应object上的权限。
	Privileges *[]string `json:"privileges,omitempty"`

	// 拥有该权限的用户名。
	UserName *string `json:"user_name,omitempty"`
}

func (o TableUserPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableUserPrivilege struct{}"
	}

	return strings.Join([]string{"TableUserPrivilege", string(data)}, " ")
}
