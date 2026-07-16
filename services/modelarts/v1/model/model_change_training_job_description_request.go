package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeTrainingJobDescriptionRequest Request Object
type ChangeTrainingJobDescriptionRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	Body *JobDescription `json:"body,omitempty"`
}

func (o ChangeTrainingJobDescriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTrainingJobDescriptionRequest struct{}"
	}

	return strings.Join([]string{"ChangeTrainingJobDescriptionRequest", string(data)}, " ")
}
