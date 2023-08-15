package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowDto 流程请求体
type WorkflowDto struct {

	// 流程名称，取值范围[1,56]，允许大小写字母、数字、以及特殊字符中划线(-)和下划线(_)。更新流程时，流程名称不支持修改。
	Name string `json:"name"`

	// 流程版本，取值范围[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。更新流程时，流程版本不支持修改。
	Version string `json:"version"`

	// 流程简述 取值范围[0,128]
	Summary *string `json:"summary,omitempty"`

	// 流程描述 取值范围[0,65535]，后续支持markdown文本
	Description *string `json:"description,omitempty"`

	// 流程标签，取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者结尾。
	Labels *[]string `json:"labels,omitempty"`

	// 流程超时时间，取值范围[1,144000]，单位分钟，默认1440
	Timeout *int32 `json:"timeout,omitempty"`

	// 流程的当前工作目录，默认为根目录，用户可显式指定;输出路径必须以斜杠（/）开头且不能以斜杠（/）结尾，不能包含两个以上相邻的斜杠（/），不能包含以下特殊字符：\\ : ; * ? < \" > | 。其中单个文件夹名称不能以中划线（-）开头，不能以英文句号（.）或斜杠（/）或空格开头或结尾
	OutputDir *string `json:"output_dir,omitempty"`

	// 流程中子任务的描述信息，子任务数量取值范围:[1,64]
	Tasks *[]TaskDto `json:"tasks,omitempty"`
}

func (o WorkflowDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowDto struct{}"
	}

	return strings.Join([]string{"WorkflowDto", string(data)}, " ")
}
