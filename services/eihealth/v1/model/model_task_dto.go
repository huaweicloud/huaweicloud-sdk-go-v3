package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskDto struct {

	// 子任务实际名称，取值范围[1,32]，只能以大小写字母开头，由大小写字母、数字、中划线（-）、下划线（_）组成，以大小写字符或数字结尾。
	TaskName string `json:"task_name"`

	// 应用id，取值范围[1,135]，正则先不能有中文，两种格式。特殊id，采用{app_name}::{app_version}::{src_project_name}格式，用于手动创建场景；其他场景，app_id为系统分配的唯一标识
	AppId string `json:"app_id"`

	// 流程的子任务展示名称，最大长度64
	DisplayName *string `json:"display_name,omitempty"`

	// 子任务的输出存放路径，用户可显式指定;输出路径必须以斜杠（/）开头且不能以斜杠（/）结尾，不能包含两个以上相邻的斜杠（/），不能包含以下特殊字符：\\ : ; * ? < \" > | 。其中单个文件夹名称不能以中划线（-）开头，不能以英文句号（.）或斜杠（/）或空格开头或结尾
	OutputDir *string `json:"output_dir,omitempty"`

	Resources *TaskResourceDto `json:"resources,omitempty"`

	Location *VertexLocationDto `json:"location,omitempty"`

	// 任务的输入参数信息
	Inputs *[]TaskParameterDto `json:"inputs,omitempty"`
}

func (o TaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDto struct{}"
	}

	return strings.Join([]string{"TaskDto", string(data)}, " ")
}
