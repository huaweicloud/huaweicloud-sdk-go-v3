package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateImageTranslateRequestBody struct {
	Input *ImageTranslateTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *ImageTranslateConfig `json:"config"`
}

func (o CreateImageTranslateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageTranslateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageTranslateRequestBody", string(data)}, " ")
}
