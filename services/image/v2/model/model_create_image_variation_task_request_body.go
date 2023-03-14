package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateImageVariationTaskRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *ImageVariationTaskConfig `json:"config"`
}

func (o CreateImageVariationTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageVariationTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageVariationTaskRequestBody", string(data)}, " ")
}
