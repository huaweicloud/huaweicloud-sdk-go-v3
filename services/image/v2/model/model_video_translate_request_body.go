package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoTranslateRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoTranslateConfig `json:"config,omitempty"`
}

func (o VideoTranslateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateRequestBody struct{}"
	}

	return strings.Join([]string{"VideoTranslateRequestBody", string(data)}, " ")
}
