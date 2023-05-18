package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 备份里的库表信息。 - 为空表示实例级查询 - 非空表示库表级查询
type QueryDatabaseTableInfo struct {

	// 数据库名称。
	DatabaseName string `json:"database_name"`

	// 表名称列表。 - table_names为空的时候，表示库级别查询。 - table_names非空的时候，表示表级别查询。
	TableNames *[]string `json:"table_names,omitempty"`
}

func (o QueryDatabaseTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDatabaseTableInfo struct{}"
	}

	return strings.Join([]string{"QueryDatabaseTableInfo", string(data)}, " ")
}
