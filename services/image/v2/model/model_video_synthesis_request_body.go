package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoSynthesisRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoSynthesisConfig `json:"config"`
}

func (o VideoSynthesisRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSynthesisRequestBody struct{}"
	}

	return strings.Join([]string{"VideoSynthesisRequestBody", string(data)}, " ")
}
