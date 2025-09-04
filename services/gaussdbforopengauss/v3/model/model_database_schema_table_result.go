package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseSchemaTableResult 数据库表结构及其所属 schema
type DatabaseSchemaTableResult struct {

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`
}

func (o DatabaseSchemaTableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseSchemaTableResult struct{}"
	}

	return strings.Join([]string{"DatabaseSchemaTableResult", string(data)}, " ")
}
