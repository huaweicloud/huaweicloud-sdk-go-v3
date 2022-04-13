package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Column struct {
	// 列名称。

	ColumnName string `json:"column_name"`
	// 列描述。

	Description *string `json:"description,omitempty"`
	// 列类型。

	Type string `json:"type"`
	// 是否分区列。

	IsPartitionColumn *bool `json:"is_partition_column,omitempty"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}
