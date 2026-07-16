package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadListStatisticsStatistics **参数描述**：作业的统计信息。
type WorkloadListStatisticsStatistics struct {

	// **参数描述**：统计信息列表数量。 **取值范围**：不涉及。
	Total int32 `json:"total"`

	// **参数描述**：特定作业类型统计信息。
	Items []WorkloadStatistics `json:"items"`
}

func (o WorkloadListStatisticsStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadListStatisticsStatistics struct{}"
	}

	return strings.Join([]string{"WorkloadListStatisticsStatistics", string(data)}, " ")
}
