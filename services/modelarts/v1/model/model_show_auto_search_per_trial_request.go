package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchPerTrialRequest Request Object
type ShowAutoSearchPerTrialRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 超参搜索的trial_id。
	TrialId string `json:"trial_id"`
}

func (o ShowAutoSearchPerTrialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchPerTrialRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchPerTrialRequest", string(data)}, " ")
}
