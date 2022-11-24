package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobTaskDto struct {

	// 子任务实际名称，取值范围[1,32]，只能以大小写字母开头，由大小写字母、数字、中划线（-）、下划线（_）组成，以大小写字符或数字结尾。需要和已有子任务的名称一致。
	TaskName string `json:"task_name"`

	// 任务的输入参数信息
	Inputs *[]TaskParameterDto `json:"inputs,omitempty"`

	Resources *TaskResourceDto `json:"resources,omitempty"`
}

func (o JobTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobTaskDto struct{}"
	}

	return strings.Join([]string{"JobTaskDto", string(data)}, " ")
}
