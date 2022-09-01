package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Column struct {

	// 列名称。
	ColumnName string `json:"column_name" xml:"column_name"`

	// 列描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 列类型。
	Type string `json:"type" xml:"type"`

	// 是否分区列。
	IsPartitionColumn *bool `json:"is_partition_column,omitempty" xml:"is_partition_column"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}
