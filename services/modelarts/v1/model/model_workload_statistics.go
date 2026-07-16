package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatistics 特定作业类型统计信息。
type WorkloadStatistics struct {

	// **参数描述**： 作业类型。 **取值范围**： 可选值如下： - train：训练作业 - infer：推理作业 - notebook：Notebook作业
	Type string `json:"type"`

	// **参数描述**： 作业个数。 **取值范围**： 不涉及。
	Total int32 `json:"total"`

	Status *WorkloadStatisticsStatus `json:"status"`
}

func (o WorkloadStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatistics struct{}"
	}

	return strings.Join([]string{"WorkloadStatistics", string(data)}, " ")
}
