package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpdProjectIssueRequestBody struct {

	// 工作项名称
	Title *string `json:"title,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 状态[\"Committed\", \"Analyse\", \"ToBeConfirmed\", \"Plan\", \"Doing\", \"Delivered\", \"Checking\"]
	Status *string `json:"status,omitempty"`

	// 提出项目domainId
	SrcDomain *string `json:"src_domain,omitempty"`

	// 提交人Id
	SubmittedBy *string `json:"submitted_by,omitempty"`

	// 归属项目domainId
	DomainId *string `json:"domain_id,omitempty"`

	// 承接人id
	Recipient *[]string `json:"recipient,omitempty"`

	// 期望完成时间
	ExpectDeliveryTime *int64 `json:"expect_delivery_time,omitempty"`

	// 优先级
	Priority *string `json:"priority,omitempty"`

	// 抄送人id
	AssignedCc *[]string `json:"assigned_cc,omitempty"`

	// 工作项分类：[Epic,FE,IR,RR,SR,US,AR,Bug,Task]
	Category *string `json:"category,omitempty"`

	// 责任人
	Assignee *string `json:"assignee,omitempty"`

	// PI ID
	PlanPi *string `json:"plan_pi,omitempty"`

	// 迭代ID
	PlanIteration *string `json:"plan_iteration,omitempty"`

	// 计划开始时间
	PlanStartDate *int64 `json:"plan_start_date,omitempty"`

	// 计划结束时间
	PlanEndDate *int64 `json:"plan_end_date,omitempty"`

	// 计划工时
	WorkloadManDay *int32 `json:"workload_man_day,omitempty"`

	// 领域
	BusinessDomain *string `json:"business_domain,omitempty"`

	// 是否需要分解
	NeedBreak *string `json:"need_break,omitempty"`
}

func (o CreateIpdProjectIssueRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueRequestBody", string(data)}, " ")
}
