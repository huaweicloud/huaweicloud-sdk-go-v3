package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobRoutePlanRequest Request Object
type ShowTrainingJobRoutePlanRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowTrainingJobRoutePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobRoutePlanRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobRoutePlanRequest", string(data)}, " ")
}
