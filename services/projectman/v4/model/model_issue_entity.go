package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueEntity 工作项的公共父类，即所有工作项类型都有这些字段 其子类有 IssueAREntity IssueBugEntity IssueEntity IssueEpicEntity IssueFEEntity IssueIREntity IssueRREntity IssueSFEntity IssueSREntity IssueTaskEntity IssueUSEntity
type IssueEntity struct {

	// 工作项id
	Id *string `json:"id,omitempty"`

	// 工作项标题
	Title *string `json:"title,omitempty"`

	// 工作项描述字段
	Description *string `json:"description,omitempty"`

	// 工作项大分类定义 requirement(研发需求)、bug(缺陷)、task(任务)、feature(特性)、raw_requirement(原始需求)
	Type *string `json:"type,omitempty"`

	// 工作项编号
	Number *string `json:"number,omitempty"`

	// 工作项类型，编辑工作项时，此字段必填、值为当前工作项正确的工作项类型，但不会更新此字段
	Category *string `json:"category,omitempty"`

	// 父工作项id
	ParentId *string `json:"parent_id,omitempty"`

	// 工作项所属的项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 工作项状态code
	Status *string `json:"status,omitempty"`

	// 工作项的生命周期，可选值为\"正在工作\",\"作废\"
	State *string `json:"state,omitempty"`

	Assignee *UserEntity `json:"assignee,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	// 工作项创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	ModifiedBy *UserEntity `json:"modified_by,omitempty"`

	// 工作项最近更新时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 工作项计划结束日期，时间戳
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 工作项关闭时间
	CloseTime *string `json:"close_time,omitempty"`

	// 工作项计划工时，保留一位小数，取值范围为0~999999999.9
	Workload *string `json:"workload,omitempty"`

	// 工作项计划工时，保留一位小数，取值范围为0~999999999.9，不可编辑
	WorkloadSum *string `json:"workload_sum,omitempty"`

	// 工作项所属租户id
	TenantId *string `json:"tenant_id,omitempty"`

	// 工作项关联项id，多个关联项用英文逗号分隔，同一工作项最多支持50个关联项
	Link *string `json:"link,omitempty"`

	// 工作项是否已挂起
	Suspended *bool `json:"suspended,omitempty"`

	// 工作项状态改变时间，可用于计算工作项在当前状态停留天数
	StatusModifiedTime *string `json:"status_modified_time,omitempty"`

	// 工作项标签
	Labels *[]LabelEntity `json:"labels,omitempty"`

	// 工作项自定义字段映射，用户添加的系统字段也在此列 { \"code\":\"字段code\", \"value\":\"字段值\" }
	CustomFields *[]FieldCodeValuePair `json:"custom_fields,omitempty"`

	// 工作项的子工作项集合
	Children *[]IssueEntity `json:"children,omitempty"`

	// 子工作项的路径
	Path *string `json:"path,omitempty"`

	// IR和FE的关联字段，工作项类型为IR时，有此字段
	Ir2feature *string `json:"ir2feature,omitempty"`

	// 工作项是否需要分解,仅可以分解的工作项类型有此字段
	NeedBreak *string `json:"need_break,omitempty"`

	// 分解状态 已分解—decomposed 未分解—undecomposed 不涉及— --
	BreakStatus *string `json:"break_status,omitempty"`

	// 工作项基线状态 未基线 —— null 已基线 —— baselined 基线评审中——baseline-reviewing
	Baseline *string `json:"baseline,omitempty"`

	// 工作项优先级，部分工作项有此字段
	Priority *string `json:"priority,omitempty"`

	// 是否涉及网络安全。预设字段中，仅研发需求有此字段
	RelatedNetworkSecurity *string `json:"related_network_security,omitempty"`

	// 研发需求协同信息，协同任务id
	Collaboratives *string `json:"collaboratives,omitempty"`

	// 领域字段
	BusinessDomain *string `json:"business_domain,omitempty"`

	// 工作项发布(老版本名为PI) id
	PlanPi *string `json:"plan_pi,omitempty"`

	// 工作项变更状态 变更评审中——change-reviewing 已变更——changed 未变更-unchange或null
	ChangeStatus *string `json:"change_status,omitempty"`

	// 无需分解原因，need_break=no时有此字段
	NoBreakReason *string `json:"no_break_reason,omitempty"`

	// 提出人，部分工作项有此字段，多人时用英文逗号分隔
	SubmittedBy *string `json:"submitted_by,omitempty"`

	// IR关联的RR id，多选时用英文逗号分隔
	Ir2rr *string `json:"ir2rr,omitempty"`
}

func (o IssueEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueEntity struct{}"
	}

	return strings.Join([]string{"IssueEntity", string(data)}, " ")
}
