package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowExecutionBrief 函数流执行概要信息
type WorkflowExecutionBrief struct {

	// 流程定义ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 流程执行实例ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态
	Status *string `json:"status,omitempty"`

	// 流程实例创建时间，格式：UTC时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：UTC时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：UTC时间戳
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 流程实例创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 审批人列表
	ApproveUserNameList *[]string `json:"approve_user_name_list,omitempty"`

	// 执行记录
	ExecutionResultList *[]ExecutionResultList `json:"execution_result_list,omitempty"`

	// 租户从IAM申请到的projectid，一般为32位字符串。
	ProjectId *string `json:"project_id,omitempty"`

	// 执行工作流的修改时间，格式：UTC时间戳
	WorkflowEditTime *int64 `json:"workflow_edit_time,omitempty"`

	// 执行快照
	LastRecordIdWithSnapshot *string `json:"last_record_id_with_snapshot,omitempty"`
}

func (o WorkflowExecutionBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowExecutionBrief struct{}"
	}

	return strings.Join([]string{"WorkflowExecutionBrief", string(data)}, " ")
}
