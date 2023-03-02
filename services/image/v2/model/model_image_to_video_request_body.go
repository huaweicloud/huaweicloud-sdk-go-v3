package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageToVideoRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *ImageToVideoConfig `json:"config"`
}

func (o ImageToVideoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageToVideoRequestBody struct{}"
	}

	return strings.Join([]string{"ImageToVideoRequestBody", string(data)}, " ")
}
