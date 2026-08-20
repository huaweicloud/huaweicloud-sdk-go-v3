package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReviewEntity struct {

	// 评审单ID。
	Id *string `json:"id,omitempty"`

	// 评审单编号。
	Number *string `json:"number,omitempty"`

	// 评审单的生命周期。
	State *string `json:"state,omitempty"`

	// 评审单标题。
	Title *string `json:"title,omitempty"`

	// 评审单类别。
	Category *string `json:"category,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	ModifiedBy *UserEntity `json:"modified_by,omitempty"`

	// 评审单抄送人。
	AssignedCc *[]UserEntity `json:"assigned_cc,omitempty"`

	// 评审单创建时间戳。
	CreatedTime *string `json:"created_time,omitempty"`

	// 评审单最后修改时间戳。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 计划完成日期时间戳。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 计划开始日期时间戳。
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// 评审单完成时间。
	CloseTime *string `json:"close_time,omitempty"`

	Status *StatusEntity `json:"status,omitempty"`

	// 评审单描述。
	Description *string `json:"description,omitempty"`

	// 评审单完成时间。
	ClosedTime *string `json:"closed_time,omitempty"`

	// 决策人ID。
	Approver *string `json:"approver,omitempty"`

	// 评审专家ID。
	Reviewer *string `json:"reviewer,omitempty"`

	// 评审对象列表。
	Cos *[]CoEntity `json:"cos,omitempty"`

	// 审批信息列表。
	Ccbs *[]CcbEntity `json:"ccbs,omitempty"`

	OldStatus *StatusEntity `json:"old_status,omitempty"`

	// 抄送人列表。
	Cc *[]UserEntity `json:"cc,omitempty"`
}

func (o ReviewEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewEntity struct{}"
	}

	return strings.Join([]string{"ReviewEntity", string(data)}, " ")
}
