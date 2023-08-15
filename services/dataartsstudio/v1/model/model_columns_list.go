package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnsList schema信息
type ColumnsList struct {

	// 字段注解
	Comment *string `json:"comment,omitempty"`

	// 字段名称
	ColumnName *string `json:"column_name,omitempty"`

	// 字段类型
	ColumnType *string `json:"column_type,omitempty"`

	// 字段的顺序
	SeqNumber *int32 `json:"seq_number,omitempty"`

	// 字段是否为主键
	Primary *bool `json:"primary,omitempty"`

	// 是否对字段进行分割
	PartitionCol *bool `json:"partition_col,omitempty"`
}

func (o ColumnsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnsList struct{}"
	}

	return strings.Join([]string{"ColumnsList", string(data)}, " ")
}
