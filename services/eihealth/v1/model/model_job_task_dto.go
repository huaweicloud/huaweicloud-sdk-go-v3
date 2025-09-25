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

	// **参数解释**： 任务的输出参数信息。 **约束限制**： 最多支持128个参数。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Outputs *[]TaskParameterDto `json:"outputs,omitempty"`

	// **参数解释**： 子任务结果存储目录，默认为空。 **约束限制**： 不涉及 **取值范围**： 长度[0,128]。 **默认取值**： 不涉及
	OutputDir *string `json:"output_dir,omitempty"`

	Resources *TaskResourceDto `json:"resources,omitempty"`

	// 子任务使用的IO加速实例类型，不填表示不使用；
	IoAccType *string `json:"io_acc_type,omitempty"`
}

func (o JobTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobTaskDto struct{}"
	}

	return strings.Join([]string{"JobTaskDto", string(data)}, " ")
}
