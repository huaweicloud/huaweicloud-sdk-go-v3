package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainJobTagsRequest Request Object
type DeleteTrainJobTagsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	Body *DeleteTmsTagsRequest `json:"body,omitempty"`
}

func (o DeleteTrainJobTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainJobTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrainJobTagsRequest", string(data)}, " ")
}
