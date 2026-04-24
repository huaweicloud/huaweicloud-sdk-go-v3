package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFactoryWorkspacesApprovalRespJobApplySearchList struct {

	// 当前审批人。
	ActualApprover *string `json:"actual_approver,omitempty"`

	// 审批单ID。
	ApplyId *string `json:"apply_id,omitempty"`

	// 审核信息。
	ApprovalMsg *string `json:"approval_msg,omitempty"`

	// 审批时间。
	ApprovalTime *int32 `json:"approval_time,omitempty"`

	// 审批人。
	ApproverName *string `json:"approver_name,omitempty"`

	// 作业或脚本变更类型：修改或者删除。
	ChangeType *string `json:"change_type,omitempty"`

	// 申请人。
	CreateUser *string `json:"create_user,omitempty"`

	// 审批对象ID。
	ObjectId *string `json:"object_id,omitempty"`

	// 作业或者脚本名称。
	ObjectName *string `json:"object_name,omitempty"`

	// 审批对象类型：作业或者脚本。
	ObjectType *string `json:"object_type,omitempty"`

	// 审批状态。
	Status *string `json:"status,omitempty"`

	// 审批提交时间。
	SubmitTime *int64 `json:"submit_time,omitempty"`
}

func (o ListFactoryWorkspacesApprovalRespJobApplySearchList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryWorkspacesApprovalRespJobApplySearchList struct{}"
	}

	return strings.Join([]string{"ListFactoryWorkspacesApprovalRespJobApplySearchList", string(data)}, " ")
}
