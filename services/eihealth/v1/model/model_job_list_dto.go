package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobListDto struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name *string `json:"name,omitempty"`

	// 作业的描述,取值范围：输入字符最大长度为255
	Description *string `json:"description,omitempty"`

	// 作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 作业优先级，[0,9]，0表示最低，默认0
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行超时时长，取值范围: [1, 144000]，单位：分钟，默认数值1440
	Timeout *int32 `json:"timeout,omitempty"`

	// job结果存储目录，不指定则在workflow的工作目录下生产job同名子目录，指定则已指定路径为准
	OutputDir *string `json:"output_dir,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 失败提示，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// 失败原因，当作业执行失败时会返回
	FailedReason *string `json:"failed_reason,omitempty"`

	// 创建任务的用户名称
	UserName *string `json:"user_name,omitempty"`

	ToolInfo *ToolInfoDto `json:"tool_info,omitempty"`

	// IO加速实例ID
	IoAccId *string `json:"io_acc_id,omitempty"`

	// 作业使用的SFS-Turbo实例预期占用存储量，单位G，用于投递作业时评估当前加速实例余量是否充足
	IoAccExpectedUsage *int32 `json:"io_acc_expected_usage,omitempty"`

	// 仍在运行中的子任务
	StillRunningTasks *[]string `json:"still_running_tasks,omitempty"`
}

func (o JobListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobListDto struct{}"
	}

	return strings.Join([]string{"JobListDto", string(data)}, " ")
}
