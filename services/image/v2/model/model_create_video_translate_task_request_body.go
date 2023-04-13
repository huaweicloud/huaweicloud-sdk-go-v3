package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoTranslateTaskRequestBody struct {
	Input *VideoTranslateTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoTranslateConfig `json:"config,omitempty"`
}

func (o CreateVideoTranslateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoTranslateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoTranslateTaskRequestBody", string(data)}, " ")
}
