package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionColumnStatistics 分区中的列统计
type PartitionColumnStatistics struct {
	ColumnStatisticsDesc *PartitionColumnStatisticsDescription `json:"column_statistics_desc"`

	// 列统计信息
	ColumnStatisticsObjects []ColumnStatisticsObj `json:"column_statistics_objects"`
}

func (o PartitionColumnStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionColumnStatistics struct{}"
	}

	return strings.Join([]string{"PartitionColumnStatistics", string(data)}, " ")
}
