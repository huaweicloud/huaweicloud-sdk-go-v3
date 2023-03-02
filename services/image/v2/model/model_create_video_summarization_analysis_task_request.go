package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVideoSummarizationAnalysisTaskRequest struct {
	Body *VideoSummarizationCreateTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoSummarizationAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSummarizationAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoSummarizationAnalysisTaskRequest", string(data)}, " ")
}
