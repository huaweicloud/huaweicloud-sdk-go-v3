package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户权限信息
type DeleteDatabasePermission struct {

	// 数据库用户名
	Name string `json:"name"`

	// 主机地址
	Host *string `json:"host,omitempty"`

	// 数据库列表
	Databases []string `json:"databases"`
}

func (o DeleteDatabasePermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabasePermission struct{}"
	}

	return strings.Join([]string{"DeleteDatabasePermission", string(data)}, " ")
}
