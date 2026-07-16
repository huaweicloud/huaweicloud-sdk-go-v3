package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchTrialEarlyStopRequest Request Object
type ShowAutoSearchTrialEarlyStopRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 超参搜索的trial_id。
	TrialId string `json:"trial_id"`
}

func (o ShowAutoSearchTrialEarlyStopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchTrialEarlyStopRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchTrialEarlyStopRequest", string(data)}, " ")
}
