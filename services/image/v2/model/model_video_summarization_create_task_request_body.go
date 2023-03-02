package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoSummarizationCreateTaskRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *SummarizationAnalysisConfig `json:"config"`
}

func (o VideoSummarizationCreateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSummarizationCreateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"VideoSummarizationCreateTaskRequestBody", string(data)}, " ")
}
