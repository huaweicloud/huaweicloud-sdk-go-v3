package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建作业的请求体
type JobDto struct {

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name string `json:"name"`

	// 作业的描述,取值范围：输入字符最大长度为255
	Description *string `json:"description,omitempty"`

	// 作业标签，取值范围[0,5]，单个标签最大长度32字符，仅仅包含小写字母或数字或大写字母
	Labels *[]string `json:"labels,omitempty"`

	// 作业的优先级,取值范围[0,9]，0最低，默认数值0
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行超时时长，取值范围: [1, 144000]，单位：分钟，默认数值1440
	Timeout *int32 `json:"timeout,omitempty"`

	// 作业结果存储目录，不指定则在workflow的工作目录下生产job同名子目录，指定则已指定路径为准;输出路径必须以斜杠（/）开头且不能以斜杠（/）结尾，不能包含两个以上相邻的斜杠（/），不能包含以下特殊字符：\\ : ; * ? < \" > | 。其中单个文件夹名称不能以中划线（-）开头，不能以英文句号（.）或斜杠（/）或空格开头或结尾
	OutputDir *string `json:"output_dir,omitempty"`

	// 作业依赖的组件id，组件当前仅支持流程，取值范围[1,135]，支持大小写字母和数字。目前支持两种格式，特殊id：{流程名称}::{流程版本}::{源项目名称}；正常id：流程id
	ToolId string `json:"tool_id"`

	// 作业依赖的组件类型，仅支持填写workflow
	ToolType string `json:"tool_type"`

	// 基于替换规则压扁后，job实际的运行信息
	Tasks *[]JobTaskDto `json:"tasks,omitempty"`

	// 作业使用的SFS-Turbo实例id，不填表示不使用
	IoAccId *string `json:"io_acc_id,omitempty"`

	// 作业使用的SFS-Turbo实例预期占用存储量，单位G，用于投递作业时评估当前加速实例余量是否充足
	IoAccExpectedUsage *int32 `json:"io_acc_expected_usage,omitempty"`

	// 节点标签 取值范围[0,1]，单个标签最大长度63字符
	NodeLabels *[]string `json:"node_labels,omitempty"`
}

func (o JobDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDto struct{}"
	}

	return strings.Join([]string{"JobDto", string(data)}, " ")
}
