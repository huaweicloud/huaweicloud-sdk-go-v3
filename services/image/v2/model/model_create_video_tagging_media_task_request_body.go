package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoTaggingMediaTaskRequestBody struct {
	Input *VideoTaggingMediaTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoTaggingTaskConfig `json:"config,omitempty"`
}

func (o CreateVideoTaggingMediaTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoTaggingMediaTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoTaggingMediaTaskRequestBody", string(data)}, " ")
}
