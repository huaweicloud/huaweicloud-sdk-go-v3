package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseTableConfigCheckResult 表配置校验结果。
type ChDatabaseTableConfigCheckResult struct {

	// 数据库表名。
	TableName string `json:"table_name"`

	// 表配置项。  允许输入的列操作有：PARTITION BY, COLUMNS, ORDER BY, SAMPLE BY, PRIMARY KEY, TTL
	TableConfig string `json:"table_config"`

	// 校验结果。
	CheckResult string `json:"check_result"`
}

func (o ChDatabaseTableConfigCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseTableConfigCheckResult struct{}"
	}

	return strings.Join([]string{"ChDatabaseTableConfigCheckResult", string(data)}, " ")
}
