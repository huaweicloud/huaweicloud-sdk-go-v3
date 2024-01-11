package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnDetails 字段详细信息
type ColumnDetails struct {

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 逻辑库名称
	Schema *string `json:"schema,omitempty"`

	// 表名称
	Table *string `json:"table,omitempty"`

	// 字段名称
	Column *string `json:"column,omitempty"`
}

func (o ColumnDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnDetails struct{}"
	}

	return strings.Join([]string{"ColumnDetails", string(data)}, " ")
}
