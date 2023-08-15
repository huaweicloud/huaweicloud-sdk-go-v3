package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingRequestBody struct {
	Input *ImageHighresolutionMattingInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *ImageHighresolutionMattingConfig `json:"config"`
}

func (o ImageHighresolutionMattingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingRequestBody struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingRequestBody", string(data)}, " ")
}
