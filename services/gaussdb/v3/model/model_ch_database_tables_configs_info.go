package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseTablesConfigsInfo 表配置。
type ChDatabaseTablesConfigsInfo struct {

	// 数据库表名。
	TableName string `json:"table_name"`

	// 表配置值。  允许输入的列操作有：PARTITION BY, COLUMNS, ORDER BY, SAMPLE BY, PRIMARY KEY, TTL
	TableConfig string `json:"table_config"`
}

func (o ChDatabaseTablesConfigsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseTablesConfigsInfo struct{}"
	}

	return strings.Join([]string{"ChDatabaseTablesConfigsInfo", string(data)}, " ")
}
