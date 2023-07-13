package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskDefinitionDto struct {

	// 子任务实际名称
	TaskName *string `json:"task_name,omitempty"`

	// 流程的子任务展示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 子任务的输出存放路径
	OutputDir *string `json:"output_dir,omitempty"`

	// 子任务的完整输出路径，查看流程不会返回，查看作业时才会返回完整输出路径
	WholeOutputDir *string `json:"whole_output_dir,omitempty"`

	// 子任务使用的IO加速类型，不填表示不使用；
	IoAccType *string `json:"io_acc_type,omitempty"`

	Resources *TaskResourceDto `json:"resources,omitempty"`

	Location *VertexLocationDto `json:"location,omitempty"`

	// 子任务的输入参数信息
	Inputs *[]TaskParameterDto `json:"inputs,omitempty"`

	AppInfo *AppInfoDto `json:"app_info,omitempty"`
}

func (o TaskDefinitionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDefinitionDto struct{}"
	}

	return strings.Join([]string{"TaskDefinitionDto", string(data)}, " ")
}
