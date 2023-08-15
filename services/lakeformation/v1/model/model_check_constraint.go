package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckConstraint struct {

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 列名
	ColumnName string `json:"column_name"`

	// constraint Name
	ConstraintName string `json:"constraint_name"`

	// 检查条件表达式
	CheckExpression *string `json:"check_expression,omitempty"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o CheckConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConstraint struct{}"
	}

	return strings.Join([]string{"CheckConstraint", string(data)}, " ")
}
