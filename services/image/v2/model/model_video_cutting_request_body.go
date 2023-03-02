package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoCuttingRequestBody struct {
	Input *TaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoCuttingConfig `json:"config,omitempty"`
}

func (o VideoCuttingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCuttingRequestBody struct{}"
	}

	return strings.Join([]string{"VideoCuttingRequestBody", string(data)}, " ")
}
