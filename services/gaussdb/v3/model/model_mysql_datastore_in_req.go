package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MysqlDatastoreInReq 数据库信息。
type MysqlDatastoreInReq struct {

	// 数据库引擎，现在只支持gaussdb-mysql。
	Type string `json:"type"`

	// 数据库版本。  两位数的大版本号，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlEngineVersion.html)返回的name字段。
	Version string `json:"version"`

	// 内核数据库版本。如果需要指定具体的内核版本，请联系客服人员添加白名单。  完整的四位内核数据库版本，获取方法请参见[查询数据库引擎的版本](https://support.huaweicloud.com/api-gaussdb/ShowGaussMySqlEngineVersion.html)返回的kernel_version字段。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o MysqlDatastoreInReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlDatastoreInReq struct{}"
	}

	return strings.Join([]string{"MysqlDatastoreInReq", string(data)}, " ")
}
