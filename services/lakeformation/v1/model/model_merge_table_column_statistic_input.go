package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeTableColumnStatisticInput struct {

	// 是否是对统计信息的合并操作
	Merge *bool `json:"merge,omitempty"`

	TableColumnStatistics *TableColumnStatistics `json:"table_column_statistics"`
}

func (o MergeTableColumnStatisticInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeTableColumnStatisticInput struct{}"
	}

	return strings.Join([]string{"MergeTableColumnStatisticInput", string(data)}, " ")
}
