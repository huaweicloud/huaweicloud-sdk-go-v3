package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSqlResponseSqlOperatedObjInfo struct {

	// 列名
	ColumnName *string `json:"column_name,omitempty"`

	// 操作对象类型 - VARIABLE: 变量 - \"\": 空
	ObjectType *string `json:"object_type,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// sql类型 - LOGIN: LOGIN - CREATE_TABLE: CREATE TABLE - CREATE_TABLESPACE: CREATE TABLESPACE - DROP_TABLE: DROP TABLE - DROP_TABLESPACE: DROP TABLESPACE - DELETE: DELETE - INSERT: INSERT - INSERT_SELECT: INSERT SELECT - SELECT: SELECT - SELECT_FOR_UPDATE: SELECT FOR UPDATE - UPDATE: UPDATE - CREATE_USER: CREATE USER - DROP_USER: DROP USER - GRANT: GRANT - OPERATE: OPERATE - ALL: ALL
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
