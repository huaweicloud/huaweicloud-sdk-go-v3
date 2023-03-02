package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoSplitTaskRequestBody struct {
	Input *VideoSplitTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`
}

func (o CreateVideoSplitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoSplitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoSplitTaskRequestBody", string(data)}, " ")
}
