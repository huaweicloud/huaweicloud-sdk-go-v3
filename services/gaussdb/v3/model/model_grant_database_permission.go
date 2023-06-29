package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantDatabasePermission 数据库用户权限信息。
type GrantDatabasePermission struct {

	// 数据库用户名。
	Name string `json:"name"`

	// 主机地址。
	Host string `json:"host"`

	Databases []DatabasePermission `json:"databases"`
}

func (o GrantDatabasePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantDatabasePermission struct{}"
	}

	return strings.Join([]string{"GrantDatabasePermission", string(data)}, " ")
}
