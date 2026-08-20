package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResult 返回结果
type ProcessInstanceResponseResult struct {

	// 抄送人列表
	Cc *string `json:"cc,omitempty"`

	// 决策人ID
	Approver *string `json:"approver,omitempty"`

	// 评审单描述，列表接口不返回描述信息
	Description *string `json:"description,omitempty"`

	// 评审单完成时间
	ClosedTime *string `json:"closed_time,omitempty"`

	// 评审专家ID，逗号分隔
	Reviewer *string `json:"reviewer,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 评审单标题
	Title *string `json:"title,omitempty"`

	// 评审单最后修改时间戳
	ModifiedDate *string `json:"modified_date,omitempty"`

	CreatedBy *ProcessInstanceResponseResultCreatedBy `json:"created_by,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 评审单编号
	Number *string `json:"number,omitempty"`

	// 是否需要审批
	NeedApproval *bool `json:"need_approval,omitempty"`

	// 基线评审对象
	Br2co *string `json:"br2co,omitempty"`

	ModifiedBy *ProcessInstanceResponseResultModifiedBy `json:"modified_by,omitempty"`

	// 评审时间
	ApprovalTime *string `json:"approval_time,omitempty"`

	// 计划完成时间
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 评审单ID
	Id *string `json:"id,omitempty"`

	// 评审单工作状态，取值为\"正在工作\",\"作废\"
	State *string `json:"state,omitempty"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 类别
	Category *string `json:"category,omitempty"`

	// 计划开始时间
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	ReviewConfig *ProcessInstanceResponseResultReviewConfig `json:"review_config,omitempty"`

	Status *ProcessInstanceResponseResultStatus `json:"status,omitempty"`

	// 阶段
	Stage *string `json:"stage,omitempty"`

	// 变更对象评审专家Id列表（创建变更评审时使用）
	Opinions *[]ProcessInstanceResponseResultOpinions `json:"opinions,omitempty"`

	// 评审意见
	OpinionComments *[]string `json:"opinion_comments,omitempty"`

	// 附件
	Attachments *[]string `json:"attachments,omitempty"`

	// 关联wiki
	Wikis *[]string `json:"wikis,omitempty"`

	// 关联文档
	Associatedocuments *[]string `json:"associatedocuments,omitempty"`

	// 评审对象列表
	Cos *[]ProcessInstanceResponseResultCos `json:"cos,omitempty"`

	// 评审结果
	ApprovalPhaseResult *string `json:"approval_phase_result,omitempty"`

	// 审批信息列表
	Ccbs *[]ProcessInstanceResponseResultCcbs `json:"ccbs,omitempty"`
}

func (o ProcessInstanceResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResult struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResult", string(data)}, " ")
}
