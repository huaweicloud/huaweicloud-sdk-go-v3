package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowIpdProcessInstancesResponseResultProcessInstances struct {

	// 抄送人，多值使用英文逗号分隔。
	Cc *string `json:"cc,omitempty"`

	// 评审单决策人。
	Approver *string `json:"approver,omitempty"`

	// 评审单完成时间。
	ClosedTime *string `json:"closed_time,omitempty"`

	// 评审专家。
	Reviewer *string `json:"reviewer,omitempty"`

	// 评审分类。
	Type *string `json:"type,omitempty"`

	// 标题。
	Title *string `json:"title,omitempty"`

	// 修改时间。
	ModifiedDate *string `json:"modified_date,omitempty"`

	CreatedBy *UserVo `json:"created_by,omitempty"`

	// 项目空间ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 评审编号。
	Number *string `json:"number,omitempty"`

	// 是否需要决策人审批。
	NeedApproval *string `json:"need_approval,omitempty"`

	ModifiedBy *UserVo `json:"modified_by,omitempty"`

	// 审批时间。
	ApprovalTime *string `json:"approval_time,omitempty"`

	// 计划结束时间。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 评审单ID。
	Id *string `json:"id,omitempty"`

	// 评审单数据状态。
	State *string `json:"state,omitempty"`

	// 创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	// 评审单类型。
	Category *string `json:"category,omitempty"`

	// 计划开始时间。
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	Status *ShowIpdProcessInstancesResponseResultStatus `json:"status,omitempty"`

	// 决策人对象列表。
	Ccbs *[]UserObject `json:"ccbs,omitempty"`

	// opinion对象列表。
	Opinions *[]ShowIpdProcessInstancesResponseResultOpinions `json:"opinions,omitempty"`
}

func (o ShowIpdProcessInstancesResponseResultProcessInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProcessInstancesResponseResultProcessInstances struct{}"
	}

	return strings.Join([]string{"ShowIpdProcessInstancesResponseResultProcessInstances", string(data)}, " ")
}
