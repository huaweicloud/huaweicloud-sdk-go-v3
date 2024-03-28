package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataAuthorizationPrivilege 数据赋权的赋权信息。
type DataAuthorizationPrivilege struct {

	// 被赋权的数据对象，命名方式为： “databases.数据库名”，则数据库下面的所有数据都将被共享。 “databases.数据库名.tables.表名”, 指定的表的数据将被共享。 “databases.数据库名.tables.表名.columns.列名”，指定的列将被共享。 \"jobs.flink.flink作业ID\"，指定的作业将被共享。 \"groups.程序包组名\"，指定的程序包组将被共享。 \"resources.程序包名\"，指定程序包将被共享。
	Object string `json:"object"`

	// 待赋权、回收或更新的权限列表。 说明： 若“action”为“update”，更新列表为空，则表示回收用户在该数据库或表的所有权限
	Privileges []string `json:"privileges"`
}

func (o DataAuthorizationPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataAuthorizationPrivilege struct{}"
	}

	return strings.Join([]string{"DataAuthorizationPrivilege", string(data)}, " ")
}
