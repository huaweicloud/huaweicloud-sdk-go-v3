package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingExperimentStatistic 训练实验的统计数据。
type TrainingExperimentStatistic struct {

	// **参数解释**：当前实验下的训练作业总个数。 **取值范围**：不涉及。
	JobCount *int32 `json:"job_count,omitempty"`
}

func (o TrainingExperimentStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingExperimentStatistic struct{}"
	}

	return strings.Join([]string{"TrainingExperimentStatistic", string(data)}, " ")
}
