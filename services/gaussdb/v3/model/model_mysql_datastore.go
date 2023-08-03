package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlDatastore 数据库信息。
type MysqlDatastore struct {

	// 数据库引擎，现在只支持gaussdb-mysql
	Type string `json:"type"`

	// 数据库版本。  两位数的大版本号，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlEngineVersion.html)返回的name字段。
	Version string `json:"version"`

	// 内核数据库版本。  完整的四位内核数据库版本，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlEngineVersion.html)返回的kernel_version字段。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o MysqlDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlDatastore struct{}"
	}

	return strings.Join([]string{"MysqlDatastore", string(data)}, " ")
}
