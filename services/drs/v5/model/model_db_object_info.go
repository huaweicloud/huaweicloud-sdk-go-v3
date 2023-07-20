package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbObjectInfo 对象信息。
type DbObjectInfo struct {

	// 源数据库库名。
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 源数据库模式名。
	SourceSchemaName *string `json:"source_schema_name,omitempty"`

	// 源数据库表名。
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 目标数据库库名。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 目标数据库模式名。
	TargetSchemaName *string `json:"target_schema_name,omitempty"`

	// 目标数据库表名。
	TargetTableName *string `json:"target_table_name,omitempty"`

	// 是否有列映射。
	HasColumnInfo *bool `json:"has_column_info,omitempty"`
}

func (o DbObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObjectInfo struct{}"
	}

	return strings.Join([]string{"DbObjectInfo", string(data)}, " ")
}
