package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainJobTagsRequest Request Object
type ShowTrainJobTagsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowTrainJobTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainJobTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainJobTagsRequest", string(data)}, " ")
}
