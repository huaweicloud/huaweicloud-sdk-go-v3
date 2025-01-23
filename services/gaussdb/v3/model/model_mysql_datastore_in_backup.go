package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlDatastoreInBackup 数据库信息。
type MysqlDatastoreInBackup struct {

	// 数据库引擎，现在只支持gaussdb-mysql。
	Type string `json:"type"`

	// 数据库版本。  数据库支持的版本信息，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlEngineVersion.html)返回的name字段。
	Version string `json:"version"`
}

func (o MysqlDatastoreInBackup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlDatastoreInBackup struct{}"
	}

	return strings.Join([]string{"MysqlDatastoreInBackup", string(data)}, " ")
}
