package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaColumns 列信息
type SchemaColumns struct {

	// 是否为数组
	Array *bool `json:"array,omitempty"`

	// 列数据类型，例如 TIMESTAMP 或 INT
	ColumnDataType *string `json:"column_data_type,omitempty"`

	// 数据类型设置
	ColumnDataTypeSetting *string `json:"column_data_type_setting,omitempty"`

	// 列名称
	ColumnName *string `json:"column_name,omitempty"`

	// 列顺序号
	ColumnSequenceNumber *int32 `json:"column_sequence_number,omitempty"`

	// 列类型，例如 PHYSICAL
	ColumnType *string `json:"column_type,omitempty"`

	// 列类型设置
	ColumnTypeSetting *string `json:"column_type_setting,omitempty"`

	// 深度
	Depth *int32 `json:"depth,omitempty"`

	// 是否允许为空
	Nullable *bool `json:"nullable,omitempty"`

	// 自有名称
	OwnName *string `json:"own_name,omitempty"`

	// 父名称
	ParentName *string `json:"parent_name,omitempty"`
}

func (o SchemaColumns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaColumns struct{}"
	}

	return strings.Join([]string{"SchemaColumns", string(data)}, " ")
}
