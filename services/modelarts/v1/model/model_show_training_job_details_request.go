package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobDetailsRequest Request Object
type ShowTrainingJobDetailsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowTrainingJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobDetailsRequest", string(data)}, " ")
}
