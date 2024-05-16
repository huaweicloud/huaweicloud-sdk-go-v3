package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTableListDetail struct {

	// 库名参数必传
	DbName string `json:"db_name"`

	// schema名  备份恢复到新实例和当前实例：使用DATABASE_TABLE级别恢复参数必传, 使用DATABASE类型恢复参数无效。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名  备份恢复到新实例和当前实例：使用DATABASE_TABLE级别恢复参数必传, 使用DATABASE类型恢复参数无效。
	TableName *string `json:"table_name,omitempty"`

	// 新库名  备份恢复到新实例和当前实例：   使用DATABASE级别恢复：需注意目标实例不存在当前一样的库名，否则无法下发。   使用DATABASE_TABLE级别恢复，不填时与源库名一致。
	NewDbName *string `json:"new_db_name,omitempty"`

	// 新schema_name  备份恢复到新实例和当前实例：    使用DATABASE类型恢复参数无效。   使用DATABASE_TABLE级别恢复，不填时与源schema名一致。
	NewSchemaName *string `json:"new_schema_name,omitempty"`

	// 新表名  备份恢复到新实例和当前实例：   使用DATABASE类型恢复参数无效。   使用DATABASE_TABLE级别恢复，需注意目标实例里不存在当前的表名，否则无法下发。
	NewTableName *string `json:"new_table_name,omitempty"`
}

func (o RestoreTableListDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTableListDetail struct{}"
	}

	return strings.Join([]string{"RestoreTableListDetail", string(data)}, " ")
}
