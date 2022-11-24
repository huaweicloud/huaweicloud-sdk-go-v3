package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobResponse struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业的名称
	Name *string `json:"name,omitempty"`

	// 作业的描述
	Description *string `json:"description,omitempty"`

	// 作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 作业的优先级
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行超时时长
	Timeout *int32 `json:"timeout,omitempty"`

	// 压扁后的效果，即job运行的实际配置
	OutputDir *string `json:"output_dir,omitempty"`

	// 作业运行状态
	Status *string `json:"status,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业运行失败描述信息，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// 作业运行失败原因，当作业执行失败时会返回
	FailedReason *string `json:"failed_reason,omitempty"`

	ToolInfo *ToolInfoDto `json:"tool_info,omitempty"`

	// 基于替换规则压扁后的效果，即job运行的实际配置
	Tasks *[]TaskDefinitionDto `json:"tasks,omitempty"`

	// 作业子步骤的运行信息
	TaskRuntimeInfo *[]TaskRuntimeDto `json:"task_runtime_info,omitempty"`

	// 作业子步骤的依赖关系
	Dag map[string]map[string]string `json:"dag,omitempty"`

	// 作业使用的SFS-Turbo实例预期占用存储量，单位G，用于投递作业时评估当前加速实例余量是否充足
	IoAccExpectedUsage *int32 `json:"io_acc_expected_usage,omitempty"`

	IoAccInfo *IoAccInfoDto `json:"io_acc_info,omitempty"`

	// 计算节点标签
	NodeLabels     *[]string `json:"node_labels,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}
