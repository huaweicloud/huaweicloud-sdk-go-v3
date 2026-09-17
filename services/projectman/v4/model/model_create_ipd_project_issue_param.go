package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpdProjectIssueParam struct {

	// 工作项名称
	Title *string `json:"title,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 状态[\"Committed\", \"Analyse\", \"ToBeConfirmed\", \"Plan\", \"Doing\", \"Delivered\", \"Checking\"]
	Status *string `json:"status,omitempty"`

	// 提出项目domainId
	SrcDomain *string `json:"src_domain,omitempty"`

	// 所属特性集，适用于SF类型工作项
	FeatureSet *string `json:"feature_set,omitempty"`

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

	// 工作项层级ID
	CategoryLayerId *string `json:"category_layer_id,omitempty"`

	// 父工作项ID
	ParentId *string `json:"parent_id,omitempty"`

	// IR关联的RR的ID
	Ir2rr *string `json:"ir2rr,omitempty"`

	// US关联的RR的ID
	Us2rr *string `json:"us2rr,omitempty"`

	// 关联工作项ID，多值使用英文逗号分隔
	Link *string `json:"link,omitempty"`

	// IR关联的SF的ID
	Ir2feature *string `json:"ir2feature,omitempty"`
}

func (o CreateIpdProjectIssueParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueParam struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueParam", string(data)}, " ")
}
