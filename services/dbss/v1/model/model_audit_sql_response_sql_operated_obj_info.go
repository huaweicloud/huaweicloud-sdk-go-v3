package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSqlResponseSqlOperatedObjInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 操作对象类型
	ObjectType *string `json:"object_type,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// sql类型
	SqlType *string `json:"sql_type,omitempty"`

	// 系统名称
	SysName *string `json:"sys_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`
}

func (o AuditSqlResponseSqlOperatedObjInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlResponseSqlOperatedObjInfo struct{}"
	}

	return strings.Join([]string{"AuditSqlResponseSqlOperatedObjInfo", string(data)}, " ")
}
