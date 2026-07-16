package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowObsUrlOfTrainingJobLogsRequest Request Object
type ShowObsUrlOfTrainingJobLogsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 训练作业的任务名称。可从训练作业详情中的status.tasks字段中获取。
	TaskId string `json:"task_id"`
}

func (o ShowObsUrlOfTrainingJobLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObsUrlOfTrainingJobLogsRequest struct{}"
	}

	return strings.Join([]string{"ShowObsUrlOfTrainingJobLogsRequest", string(data)}, " ")
}
