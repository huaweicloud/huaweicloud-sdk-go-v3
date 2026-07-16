package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrainJobTagsRequest Request Object
type CreateTrainJobTagsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	Body *CreateTmsTagsRequest `json:"body,omitempty"`
}

func (o CreateTrainJobTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrainJobTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateTrainJobTagsRequest", string(data)}, " ")
}
