package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetPartitionColumnStatisticsInput 获取指定分区指定列的列统计信息
type GetPartitionColumnStatisticsInput struct {

	// 是否聚合返回统计信息
	AggregateStatics bool `json:"aggregate_statics"`

	// 统计信息的列名
	ColumnNames []string `json:"column_names"`

	// 需要统计的分区值列表
	PartitionValuesList [][]string `json:"partition_values_list"`
}

func (o GetPartitionColumnStatisticsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetPartitionColumnStatisticsInput struct{}"
	}

	return strings.Join([]string{"GetPartitionColumnStatisticsInput", string(data)}, " ")
}
