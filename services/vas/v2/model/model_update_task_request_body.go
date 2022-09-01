package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTaskRequestBody struct {

	// 作业的名称，必填。仅能包含汉字、字母、数字、中划线和下划线，长度介于1~100之间。
	Name *string `json:"name,omitempty" xml:"name"`

	// 作业的描述，选填。长度不超过500。
	Description *string `json:"description,omitempty" xml:"description"`

	Timing *TaskTiming `json:"timing,omitempty" xml:"timing"`

	Input *TaskInput `json:"input,omitempty" xml:"input"`

	Output *TaskOutput `json:"output,omitempty" xml:"output"`

	ServiceConfig *TaskServiceConfig `json:"service_config,omitempty" xml:"service_config"`
}

func (o UpdateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTaskRequestBody", string(data)}, " ")
}
