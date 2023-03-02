package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVideoSummarizationAnalysisTaskResponse struct {

	// 任务唯一标识
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVideoSummarizationAnalysisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSummarizationAnalysisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateVideoSummarizationAnalysisTaskResponse", string(data)}, " ")
}
