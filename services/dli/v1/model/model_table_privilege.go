package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TablePrivilege 查看表的用户权限响应的权限信息。
type TablePrivilege struct {

	// 该用户有权限的对象： “databases.数据库名.tables.表名”，用户在当前表上的权限。 “databases.数据库名.tables.表名.columns.列名”，用户在列上的权限。
	Object *string `json:"object,omitempty"`

	// 用户在指定对象上的权限。
	Privileges *[]string `json:"privileges,omitempty"`
}

func (o TablePrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TablePrivilege struct{}"
	}

	return strings.Join([]string{"TablePrivilege", string(data)}, " ")
}
