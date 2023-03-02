package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageTranslateRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *ImageTranslateConfig `json:"config"`
}

func (o ImageTranslateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTranslateRequestBody struct{}"
	}

	return strings.Join([]string{"ImageTranslateRequestBody", string(data)}, " ")
}
