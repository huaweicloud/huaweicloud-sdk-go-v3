package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableColumnDto 字段信息
type TableColumnDto struct {

	// 字段名称
	ColumnName *string `json:"column_name,omitempty"`

	// 字段描述
	Description *string `json:"description,omitempty"`

	// 字段类型
	Type *string `json:"type,omitempty"`

	// 是否是分区字段
	IsPartitionColumn *bool `json:"is_partition_column,omitempty"`
}

func (o TableColumnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableColumnDto struct{}"
	}

	return strings.Join([]string{"TableColumnDto", string(data)}, " ")
}
