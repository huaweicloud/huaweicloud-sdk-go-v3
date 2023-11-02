package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlDatastoreInRes 数据库信息。
type MysqlDatastoreInRes struct {

	// 数据库引擎，现在只支持gaussdb-mysql。
	Type string `json:"type"`

	// 数据库版本。  两位数的大版本号，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdbformysql/ShowGaussMySqlEngineVersion.html)返回的name字段。
	Version string `json:"version"`

	// 内核数据库版本。  完整的四位内核数据库版本号，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdbformysql/ShowGaussMySqlEngineVersion.html)返回的kernel_version字段。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o MysqlDatastoreInRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlDatastoreInRes struct{}"
	}

	return strings.Join([]string{"MysqlDatastoreInRes", string(data)}, " ")
}
