package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProcessInstanceResponseResultCos struct {

	// 区域
	Region *string `json:"region,omitempty"`

	// 变更对象工作项类型，此处固定为CO
	Category *string `json:"category,omitempty"`

	// 评审单标题
	Title *string `json:"title,omitempty"`

	// 变更对象状态
	Status *string `json:"status,omitempty"`

	Assignee *ProcessInstanceResponseResultAssignee `json:"assignee,omitempty"`

	// 评审单描述
	Description *string `json:"description,omitempty"`

	// 变更对象关联的工作项编号
	Number *string `json:"number,omitempty"`

	// 排序
	Order *string `json:"order,omitempty"`

	// 关联的变更评审标识
	Co2cr *string `json:"co2cr,omitempty"`

	// 关联的基线评审标识
	Co2br *string `json:"co2br,omitempty"`

	// 关联的通用评审标识
	Co2gr *string `json:"co2gr,omitempty"`

	// 审批对象Id
	Id *string `json:"id,omitempty"`

	// 评审单类型
	Type *string `json:"type,omitempty"`

	// 评审单工作状态，取值为\"正在工作\",\"作废\"
	State *string `json:"state,omitempty"`

	// 变更对象工作项修改前内容
	BeforeChange *string `json:"before_change,omitempty"`

	// 变更对象修改后内容
	AfterChange *string `json:"after_change,omitempty"`

	// 评审单最后修改人
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 评审单最后修改时间
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 评审单创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 评审单创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 工作项所属租户ID，可通过[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**tenant_id**字段的值就是工作项所属租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 工作项状态
	StatusMap *string `json:"status_map,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 源系统
	SourceSystem *string `json:"source_system,omitempty"`

	// 源系统链接
	SourceSystemLink *string `json:"source_system_link,omitempty"`

	// 变更对象关联的工作项类型
	IssueCategory *string `json:"issue_category,omitempty"`

	// 工作项ID
	IssueId *string `json:"issue_id,omitempty"`

	IssueStatus *ProcessInstanceResponseResultIssueStatus `json:"issue_status,omitempty"`

	// 工作项严重程度
	IssueSeverity *string `json:"issue_severity,omitempty"`

	IssuePriority *ProcessInstanceResponseResultIssuePriority `json:"issue_priority,omitempty"`

	// 归属项目名称
	DomainTitle *string `json:"domain_title,omitempty"`

	// 提出项目名称
	SrcDomainTitle *string `json:"src_domain_title,omitempty"`

	// 责任人昵称
	IssueAssigneeName *string `json:"issue_assignee_name,omitempty"`

	// 评审原因
	ChangeReason *string `json:"change_reason,omitempty"`

	// 评审类型
	ChangeType *string `json:"change_type,omitempty"`

	// 源系统id
	SourceSystemId *string `json:"source_system_id,omitempty"`

	// 评审描述
	ChangeDescription *string `json:"change_description,omitempty"`

	// 是否已删除
	HasDeleted *string `json:"has_deleted,omitempty"`

	// 评审结果
	ApprovalPhaseResult *string `json:"approval_phase_result,omitempty"`

	// 评审完成时间
	ApprovalCompleteTime *string `json:"approval_complete_time,omitempty"`

	// 评审描述
	CcbDescription *string `json:"ccb_description,omitempty"`

	// 评审专家
	ActualCcb *string `json:"actual_ccb,omitempty"`

	// 审批信息列表
	Ccbs *string `json:"ccbs,omitempty"`

	// 评审信息
	CcbInfo *string `json:"ccb_info,omitempty"`

	// 变更对象评审专家Id列表（创建变更评审时使用）
	Opinions *string `json:"opinions,omitempty"`

	// 评审意见
	OpinionComments *string `json:"opinion_comments,omitempty"`

	// 审批时间
	ApprovalTime *string `json:"approval_time,omitempty"`

	// 租户id
	SrcDomainId *string `json:"src_domain_id,omitempty"`

	// 是否跨租户
	CrossDomain *string `json:"cross_domain,omitempty"`

	// 归属项目是否迁移
	DomainMoved *string `json:"domain_moved,omitempty"`

	// 评审专家
	Reviewer *[]string `json:"reviewer,omitempty"`

	// 决策人
	Approver *[]string `json:"approver,omitempty"`

	// 评审轮次
	Rounds *string `json:"rounds,omitempty"`

	// 最近一轮决策结果
	LastRoundResult *string `json:"last_round_result,omitempty"`
}

func (o ProcessInstanceResponseResultCos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultCos struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultCos", string(data)}, " ")
}
