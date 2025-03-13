package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionApprovalOpenapiDto struct {

	// 审批外发失败消息
	ApprovalDispatchErrorMsg *string `json:"approval_dispatch_error_msg,omitempty"`

	// 审批外发状态，0表示成功，1表示失败，null表示非SMN节点
	ApprovalDispatchStatus *string `json:"approval_dispatch_status,omitempty"`

	// 申请类型, DATA_PERMISSION
	ApprovalType *PermissionApprovalOpenapiDtoApprovalType `json:"approval_type,omitempty"`

	// 申请原因
	ApproveReason *string `json:"approve_reason,omitempty"`

	// 当前审批节点id
	CurrentNodeId *string `json:"current_node_id,omitempty"`

	// 当前审批节点审批人
	CurrentNodeName *string `json:"current_node_name,omitempty"`

	// 当前审批节点审批人类型
	CurrentNodeType *string `json:"current_node_type,omitempty"`

	Detail *PermissionApprovalDetailDto `json:"detail,omitempty"`

	// 工单结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 到期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 工单id
	Id *string `json:"id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 审批人所在权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 申请人id
	ProposerId *string `json:"proposer_id,omitempty"`

	// 申请人名称
	ProposerName *string `json:"proposer_name,omitempty"`

	// 用户申请权限时所在工作空间id
	ProposerWorkspaceId *string `json:"proposer_workspace_id,omitempty"`

	// 拒绝理由
	Reason *string `json:"reason,omitempty"`

	// 工单开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 工单状态, WAITING_APPROVE,APPROVED,REJECT,REVOKE
	Status *PermissionApprovalOpenapiDtoStatus `json:"status,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工作空间名称
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

func (o PermissionApprovalOpenapiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionApprovalOpenapiDto struct{}"
	}

	return strings.Join([]string{"PermissionApprovalOpenapiDto", string(data)}, " ")
}

type PermissionApprovalOpenapiDtoApprovalType struct {
	value string
}

type PermissionApprovalOpenapiDtoApprovalTypeEnum struct {
	DATA_PERMISSION PermissionApprovalOpenapiDtoApprovalType
}

func GetPermissionApprovalOpenapiDtoApprovalTypeEnum() PermissionApprovalOpenapiDtoApprovalTypeEnum {
	return PermissionApprovalOpenapiDtoApprovalTypeEnum{
		DATA_PERMISSION: PermissionApprovalOpenapiDtoApprovalType{
			value: "DATA_PERMISSION",
		},
	}
}

func (c PermissionApprovalOpenapiDtoApprovalType) Value() string {
	return c.value
}

func (c PermissionApprovalOpenapiDtoApprovalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionApprovalOpenapiDtoApprovalType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PermissionApprovalOpenapiDtoStatus struct {
	value string
}

type PermissionApprovalOpenapiDtoStatusEnum struct {
	WAITING_APPROVE PermissionApprovalOpenapiDtoStatus
	APPROVED        PermissionApprovalOpenapiDtoStatus
	REJECT          PermissionApprovalOpenapiDtoStatus
	REVOKE          PermissionApprovalOpenapiDtoStatus
}

func GetPermissionApprovalOpenapiDtoStatusEnum() PermissionApprovalOpenapiDtoStatusEnum {
	return PermissionApprovalOpenapiDtoStatusEnum{
		WAITING_APPROVE: PermissionApprovalOpenapiDtoStatus{
			value: "WAITING_APPROVE",
		},
		APPROVED: PermissionApprovalOpenapiDtoStatus{
			value: "APPROVED",
		},
		REJECT: PermissionApprovalOpenapiDtoStatus{
			value: "REJECT",
		},
		REVOKE: PermissionApprovalOpenapiDtoStatus{
			value: "REVOKE",
		},
	}
}

func (c PermissionApprovalOpenapiDtoStatus) Value() string {
	return c.value
}

func (c PermissionApprovalOpenapiDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionApprovalOpenapiDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
