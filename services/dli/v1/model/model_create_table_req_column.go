package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableReqColumn 新增表的列的相关参数。
type CreateTableReqColumn struct {

	// 列的名称。
	ColumnName string `json:"column_name"`

	// 列的数据类型。
	Type string `json:"type"`

	// 列的描述信息。
	Description *string `json:"description,omitempty"`

	// 表示该列是否为分区列。“true”表示为分区列，“false”为非分区列，默认为“false”。
	IsPartitionColumn *bool `json:"is_partition_column,omitempty"`
}

func (o CreateTableReqColumn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableReqColumn struct{}"
	}

	return strings.Join([]string{"CreateTableReqColumn", string(data)}, " ")
}
