package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Table struct {

	// 指定数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// 指定模式名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 指定表名。
	TableName *string `json:"table_name,omitempty"`
}

func (o Table) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Table struct{}"
	}

	return strings.Join([]string{"Table", string(data)}, " ")
}
