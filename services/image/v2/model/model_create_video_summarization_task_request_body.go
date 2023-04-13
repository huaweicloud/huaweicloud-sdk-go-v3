package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoSummarizationTaskRequestBody struct {
	Input *VideoSummarizationTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *SummarizationAnalysisConfig `json:"config"`
}

func (o CreateVideoSummarizationTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSummarizationTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoSummarizationTaskRequestBody", string(data)}, " ")
}
