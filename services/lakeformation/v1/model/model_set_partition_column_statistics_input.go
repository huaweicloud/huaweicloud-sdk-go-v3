package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPartitionColumnStatisticsInput 设置分区的列统计信息body体
type SetPartitionColumnStatisticsInput struct {

	// 是否合入原有统计信息
	NeedMerge bool `json:"need_merge"`

	// 分区统计信息的统计列表
	Statistics []PartitionColumnStatistics `json:"statistics"`
}

func (o SetPartitionColumnStatisticsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPartitionColumnStatisticsInput struct{}"
	}

	return strings.Join([]string{"SetPartitionColumnStatisticsInput", string(data)}, " ")
}
