package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowJobListDto struct {

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name *string `json:"name,omitempty"`

	// 作业的描述,取值范围：输入字符最大长度为255
	Description *string `json:"description,omitempty"`

	// 作业标签
	Labels *[]string `json:"labels,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 是否包含已被忽略的失败tasks
	HasIgnoreFailedTasks *bool `json:"has_ignore_failed_tasks,omitempty"`

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

	// 流程名称
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 流程id
	WorkflowId *string `json:"workflow_id,omitempty"`
}

func (o NextflowJobListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowJobListDto struct{}"
	}

	return strings.Join([]string{"NextflowJobListDto", string(data)}, " ")
}
