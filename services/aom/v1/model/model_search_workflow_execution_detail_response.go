package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchWorkflowExecutionDetailResponse Response Object
type SearchWorkflowExecutionDetailResponse struct {

	// 流程定义ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 流程执行实例ID。
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态。
	Status *string `json:"status,omitempty"`

	// 流程实例创建时间，格式：UTC时间戳
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：UTC时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间。
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 节点执行信息。
	ExecutionResultList *[]ExecutionResultList `json:"execution_result_list,omitempty"`

	// 审批用户列表
	ApproveUserNameList *[]string `json:"approve_user_name_list,omitempty"`

	// 租户从IAM申请到的projectid，一般为32位字符串。
	ProjectId *string `json:"project_id,omitempty"`

	// 执行workflow的更新时间
	WorkflowEditTime *int64 `json:"workflow_edit_time,omitempty"`

	// 执行快照
	LastRecordIdWithSnapshot *string `json:"last_record_id_with_snapshot,omitempty"`
	HttpStatusCode           int     `json:"-"`
}

func (o SearchWorkflowExecutionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchWorkflowExecutionDetailResponse struct{}"
	}

	return strings.Join([]string{"SearchWorkflowExecutionDetailResponse", string(data)}, " ")
}
