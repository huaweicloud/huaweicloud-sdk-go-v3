package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopTrainingJobRequest Request Object
type StopTrainingJobRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	Body *JobActionType `json:"body,omitempty"`
}

func (o StopTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"StopTrainingJobRequest", string(data)}, " ")
}
