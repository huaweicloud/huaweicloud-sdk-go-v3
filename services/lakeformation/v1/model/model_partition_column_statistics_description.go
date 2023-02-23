package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 列统计信息的描述信息
type PartitionColumnStatisticsDescription struct {

	// 分区值的列表
	PartitionValues *[]string `json:"partition_values,omitempty"`

	// 最后统计时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time"`
}

func (o PartitionColumnStatisticsDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionColumnStatisticsDescription struct{}"
	}

	return strings.Join([]string{"PartitionColumnStatisticsDescription", string(data)}, " ")
}
