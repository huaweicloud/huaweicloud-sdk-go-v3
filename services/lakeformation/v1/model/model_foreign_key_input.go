package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForeignKeyInput struct {

	// 被引用表的数据库名
	ParentKeyDatabaseName string `json:"parent_key_database_name"`

	// 被引用表的表名
	ParentKeyTableName string `json:"parent_key_table_name"`

	// 被引用列名
	ParentKeyColumnName string `json:"parent_key_column_name"`

	// 被引用键名称
	ParentKeyName string `json:"parent_key_name"`

	// 引用列名
	ForeignKeyColumnName string `json:"foreign_key_column_name"`

	// 外键名称
	ForeignKeyName string `json:"foreign_key_name"`

	// 当被引用表中被引用的记录被删除，本表中对应记录的删除规则
	DeleteRule int32 `json:"delete_rule"`

	// 外键是否启用
	EnableConstraint bool `json:"enable_constraint"`

	// 外键排列规则
	KeySequence int32 `json:"key_sequence"`

	// is foreign Key rely
	RelyConstraint bool `json:"rely_constraint"`

	// 当被引用表中被引用的记录被修改，本表中对应记录的更新规则
	UpdateRule int32 `json:"update_rule"`

	// 外键是否可用
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o ForeignKeyInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForeignKeyInput struct{}"
	}

	return strings.Join([]string{"ForeignKeyInput", string(data)}, " ")
}
