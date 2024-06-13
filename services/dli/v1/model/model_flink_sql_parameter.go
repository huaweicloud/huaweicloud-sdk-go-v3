package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkSqlParameter Flink SQL 作业参数。
type FlinkSqlParameter struct {

	// Flink SQL 语句。长度限制：1048576个字符。
	SqlBody *string `json:"sql_body,omitempty"`

	// Flink SQL 作业依赖的 Jar 包所在的 OBS 路径。
	DependencyJars *string `json:"dependency_jars,omitempty"`
}

func (o FlinkSqlParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlParameter struct{}"
	}

	return strings.Join([]string{"FlinkSqlParameter", string(data)}, " ")
}
