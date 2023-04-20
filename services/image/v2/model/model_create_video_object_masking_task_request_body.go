package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoObjectMaskingTaskRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoObjectMaskingTaskConfig `json:"config"`
}

func (o CreateVideoObjectMaskingTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoObjectMaskingTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoObjectMaskingTaskRequestBody", string(data)}, " ")
}
