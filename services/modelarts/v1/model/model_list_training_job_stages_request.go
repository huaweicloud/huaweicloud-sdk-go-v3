package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobStagesRequest Request Object
type ListTrainingJobStagesRequest struct {

	// **参数解释**：训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TrainingJobId string `json:"training_job_id"`
}

func (o ListTrainingJobStagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobStagesRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingJobStagesRequest", string(data)}, " ")
}
