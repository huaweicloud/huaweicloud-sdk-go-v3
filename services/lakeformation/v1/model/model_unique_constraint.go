package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UniqueConstraint struct {

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

	// 列序号（限制条件中第几位）
	KeySequence *int32 `json:"key_sequence,omitempty"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o UniqueConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UniqueConstraint struct{}"
	}

	return strings.Join([]string{"UniqueConstraint", string(data)}, " ")
}
