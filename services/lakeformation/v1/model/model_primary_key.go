package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrimaryKey 主键信息
type PrimaryKey struct {

	// catalog名称
	CatalogName string `json:"catalog_name"`

	// 列名称
	ColumnName string `json:"column_name"`

	// 数据库名称
	DatabaseName string `json:"database_name"`

	// 表名字
	TableName string `json:"table_name"`

	// 主键名称
	PrimaryKeyName string `json:"primary_key_name"`

	// 是否启用主键
	EnableConstraint bool `json:"enable_constraint"`

	// 主键排序顺序
	KeySequence *int32 `json:"key_sequence,omitempty"`

	// 是否被外键依赖
	RelyConstraint bool `json:"rely_constraint"`

	// 限制条件是否可用
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o PrimaryKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrimaryKey struct{}"
	}

	return strings.Join([]string{"PrimaryKey", string(data)}, " ")
}
