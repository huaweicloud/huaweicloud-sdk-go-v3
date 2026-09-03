package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFtMetricsRequest Request Object
type ShowFtMetricsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowFtMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFtMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowFtMetricsRequest", string(data)}, " ")
}
