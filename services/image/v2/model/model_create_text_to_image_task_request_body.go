package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTextToImageTaskRequestBody struct {
	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *TextToImageTaskConfig `json:"config"`
}

func (o CreateTextToImageTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTextToImageTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTextToImageTaskRequestBody", string(data)}, " ")
}
