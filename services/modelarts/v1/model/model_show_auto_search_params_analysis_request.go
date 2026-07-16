package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchParamsAnalysisRequest Request Object
type ShowAutoSearchParamsAnalysisRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`
}

func (o ShowAutoSearchParamsAnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchParamsAnalysisRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchParamsAnalysisRequest", string(data)}, " ")
}
