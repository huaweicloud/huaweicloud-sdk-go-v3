package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVideoSummarizationAnalysisTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ShowVideoSummarizationAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoSummarizationAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowVideoSummarizationAnalysisTaskRequest", string(data)}, " ")
}
