package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainingJobRequest Request Object
type DeleteTrainingJobRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o DeleteTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrainingJobRequest", string(data)}, " ")
}
