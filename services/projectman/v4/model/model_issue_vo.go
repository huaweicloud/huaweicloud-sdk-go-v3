package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueVo 工作项对象
type IssueVo struct {

	// **参数解释：**  分析结论，通常在接纳RR时填写。 **取值范围：**  不涉及。
	SysAnalysisConclusion *string `json:"sys_analysis_conclusion,omitempty"`

	// **参数解释：**  备注。通常在提交验收RR时填写。 **取值范围：**  不涉及。
	SysRemark *string `json:"sys_remark,omitempty"`

	Promised *OptionVo `json:"promised,omitempty"`

	// **参数解释：**  工作项的分类。 **取值范围：**  - requirement - raw requirement - bug - task - feature
	Type *string `json:"type,omitempty"`

	// **参数解释：**  标识工作项是否跨项目提交。 **取值范围：**  - 1：跨项目提交工作项。 - 0：非跨项目提交工作项。
	BelongInside *string `json:"belong_inside,omitempty"`

	SrcDomain *DomainVo `json:"src_domain,omitempty"`

	DomainId *DomainVo `json:"domain_id,omitempty"`

	// **参数解释：**  原始需求的协同上游需求Id。 **取值范围：**  不涉及。
	SendFrom *string `json:"send_from,omitempty"`

	// **参数解释：**  工作项编号，由工作项类型+年月日+6位随机数组成。 **取值范围：**  不涉及。
	Number *string `json:"number,omitempty"`

	// **参数解释：**  原始需求的协同下游需求Id。 **取值范围：**  不涉及。
	SendTo *string `json:"send_to,omitempty"`

	// **参数解释：**  工作项父子挂载路径关系。 **取值范围：**  不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：**  工作项计划工时。 **取值范围：**  不涉及。
	WorkloadManDay *string `json:"workload_man_day,omitempty"`

	// **参数解释：**  验收结论。通常是验收RR时填写。 **取值范围：**  不涉及。
	SysCheckConclusion *string `json:"sys_check_conclusion,omitempty"`

	// **参数解释：**  工作项唯一Id。 **取值范围：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  工作项生命周期。 **取值范围：**  - 正在工作：可正常操作的工作项； - 作废：软删除后的工作项，可在回收站恢复； - 删除：彻底删除后的工作项，无法恢复。
	State *string `json:"state,omitempty"`

	// **参数解释：**  工作项在当前状态的停留天数。 **取值范围：**  不涉及。
	StayDays *int32 `json:"stay_days,omitempty"`

	// **参数解释：**  抄送人。 **取值范围：**  不涉及。
	AssignedCc *[]UserVo `json:"assigned_cc,omitempty"`

	// **参数解释：**  工作项提交时间，指工作项进入工作流的时间，而不是创建时间。 **取值范围：**  不涉及。
	SubmitTime *string `json:"submit_time,omitempty"`

	// **参数解释：**  工作项标签。 **取值范围：**  不涉及。
	Workitem2label *[]WorkItemLabelVo `json:"workitem2label,omitempty"`

	// **参数解释：**  退回原因。通常为退回RR/Bug时填写。 **取值范围：**  不涉及。
	SysReturnConclusion *string `json:"sys_return_conclusion,omitempty"`

	// **参数解释：**  工作项完成时间。 **取值范围：**  不涉及。
	CloseTime *string `json:"close_time,omitempty"`

	Priority *OptionVo `json:"priority,omitempty"`

	// **参数解释：**  工作项最新修改时间。 **取值范围：**  不涉及。
	ModifiedDate *string `json:"modified_date,omitempty"`

	CreatedBy *UserVo `json:"created_by,omitempty"`

	BreakStatus *OptionVo `json:"break_status,omitempty"`

	// **参数解释：**  工作项上一次流转状态的时间，可用于计算停留天数。unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	StatusModifiedDate *string `json:"status_modified_date,omitempty"`

	// **参数解释：**  期望完成时间。Unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	ExpectDeliveryTime *string `json:"expect_delivery_time,omitempty"`

	// **参数解释：**  工作项的父工作项Id。 **取值范围：**  不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	Assignee *UserVo `json:"assignee,omitempty"`

	// **参数解释：**  工作项所属租户的域。 **取值范围：**  不涉及。
	Region *string `json:"region,omitempty"`

	Status *AlmStatus `json:"status,omitempty"`

	// **参数解释：**  工作项所属租户Id。 **取值范围：**  不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	PlanPi *PlanVo `json:"plan_pi,omitempty"`

	// **参数解释：**  关联工作项的关系字段。多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Link *string `json:"link,omitempty"`

	// **参数解释：**  工作项描述，最多支持50w字符。 **取值范围：**  不涉及。
	Description *string `json:"description,omitempty"`

	IsSuspended *OptionVo `json:"is_suspended,omitempty"`

	ChangeStatus *OptionVo `json:"change_status,omitempty"`

	// **参数解释：**  工作项标题。 **取值范围：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  工作项实际工时。 **取值范围：**  不涉及。
	SumWorkloadManDay *string `json:"sum_workload_man_day,omitempty"`

	// **参数解释：**  工作项关闭原因。 **取值范围：**  不涉及。
	SysCloseReason *string `json:"sys_close_reason,omitempty"`

	// **参数解释：**  重新提交原因，通常用于RR/Bug退回后重新提交。 **取值范围：**  不涉及。
	SysResubmitReason *string `json:"sys_resubmit_reason,omitempty"`

	// **参数解释：**  工作项计划完成时间。 **取值范围：**  不涉及。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// **参数解释：**  RR的子IR的id。多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Rr2ir *string `json:"rr2ir,omitempty"`

	// **参数解释：**  工作项类型层级id。 **取值范围：**  不涉及。
	CategoryLayerId *string `json:"category_layer_id,omitempty"`

	// **参数解释：**  工作项提交人。 **取值范围：**  不涉及。
	SubmittedBy *[]UserVo `json:"submitted_by,omitempty"`

	// **参数解释：**  RR的子US的id，多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Rr2us *string `json:"rr2us,omitempty"`

	// **参数解释：**  工作项无需开发原因。 **取值范围：**  不涉及。
	SysNoDevelopReason *string `json:"sys_no_develop_reason,omitempty"`

	PlanIteration *PlanVo `json:"plan_iteration,omitempty"`

	// **参数解释：**  退回原因。通常用于RR/bug退回。 **取值范围：**  不涉及。
	SysReturnReason *string `json:"sys_return_reason,omitempty"`

	// **参数解释：**  是否级联删除标记。 **取值范围：**  不涉及。
	CascadeDelete *string `json:"cascade_delete,omitempty"`

	// **参数解释：**  承接人。通常用于RR。 **取值范围：**  不涉及。
	Recipient *[]UserVo `json:"recipient,omitempty"`

	ModifiedBy *UserVo `json:"modified_by,omitempty"`

	// **参数解释：**  工作项创建时间。 **取值范围：**  不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释：**  工作项类型。 **取值范围：**  - 系统设备类项目：RR/SF/IR/SR/AR/Task/Bug。 - 独立软件类项目：RR/SF/IR/US/Task/Bug。 - 云服务类项目：RR/Epic/FE/US/Task/Bug。
	Category *string `json:"category,omitempty"`

	// **参数解释：**  研发需求协同需求状态。 **取值范围：**  不涉及。
	CollaborativeStatus *[]string `json:"collaborative_status,omitempty"`

	Project *DomainVo `json:"project,omitempty"`

	// **参数解释：**  子工作项列表。 **取值范围：**  不涉及。
	ChildIssues map[string]IssueVo `json:"child_issues,omitempty"`

	// **参数解释：**  激活次数。Bug激活时自动赋值。 **取值范围：**  不涉及。
	ActivateTimes *int32 `json:"activate_times,omitempty"`

	Baseline *OptionVo `json:"baseline,omitempty"`

	BusinessDomain *OptionVo `json:"business_domain,omitempty"`

	// **参数解释：**  子工作项Id，多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Children *string `json:"children,omitempty"`

	// **参数解释：**  协同需求的状态变化历史记录，内容为Json字符串。 **取值范围：**  不涉及。
	CollaborativeHistory *string `json:"collaborative_history,omitempty"`

	// **参数解释：**  协同需求中的记录Id。 **取值范围：**  不涉及。
	Collaboratives *string `json:"collaboratives,omitempty"`

	// **参数解释：**  卷积实际工时。父工作项中将子工作项/协同工作项的实际工时卷积得到。 **取值范围：**  不涉及。
	ConvolutionActualHours *string `json:"convolution_actual_hours,omitempty"`

	// **参数解释：**  卷积计划工时。父工作项中将子工作项/协同工作项的计划工时卷积得到。 **取值范围：**  不涉及。
	ConvolutionPlanHours *string `json:"convolution_plan_hours,omitempty"`

	// **参数解释：**  开发责任人。通常用于“开发”状态节点责任人。 **取值范围：**  不涉及。
	DevelopOwner *string `json:"develop_owner,omitempty"`

	DoneRatio *OptionVo `json:"done_ratio,omitempty"`

	// **参数解释：**  期望修复时间。 **取值范围：**  不涉及。
	ExpectedRepairDate *string `json:"expected_repair_date,omitempty"`

	// **参数解释：**  SF的子IR的id。多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Feature2ir *string `json:"feature2ir,omitempty"`

	FeatureSet *OptionVo `json:"feature_set,omitempty"`

	FoundEnv *OptionVo `json:"found_env,omitempty"`

	FoundIteration *PlanVo `json:"found_iteration,omitempty"`

	FoundPi *PlanVo `json:"found_pi,omitempty"`

	// **参数解释：**  功能场景。 **取值范围：**  不涉及。
	FunctionScene *string `json:"function_scene,omitempty"`

	// **参数解释：**  IR关联的SF的Id，一个IR仅能关联一个SF。 **取值范围：**  不涉及。
	Ir2feature *string `json:"ir2feature,omitempty"`

	// **参数解释：**  IR关联父RR的Id，多值使用英文逗号分隔。 **取值范围：**  不涉及。
	Ir2rr *string `json:"ir2rr,omitempty"`

	// **参数解释：**  工作项关联的决策意见Id。 **取值范围：**  不涉及。
	IssueOpinionId *string `json:"issue_opinion_id,omitempty"`

	// **参数解释：**  工作项关联的评审意见Id。 **取值范围：**  不涉及。
	IssueReviewId *string `json:"issue_review_id,omitempty"`

	Module *OptionVo `json:"module,omitempty"`

	NeedBreak *OptionVo `json:"need_break,omitempty"`

	NeedDevelop *OptionVo `json:"need_develop,omitempty"`

	// **参数解释：**  无需分解原因。 **取值范围：**  不涉及。
	NoBreakReason *string `json:"no_break_reason,omitempty"`

	// **参数解释：**  无需开发原因。 **取值范围：**  不涉及。
	NoDevelopReason *string `json:"no_develop_reason,omitempty"`

	// **参数解释：**  优先级顺序。 **取值范围：**  1~100。
	Order *string `json:"order,omitempty"`

	// **参数解释：**  计划开发结束时间。通常用于“开发”状态节点，Unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	PlanDevEndDate *string `json:"plan_dev_end_date,omitempty"`

	// **参数解释：**  计划处理中结束时间。通常用于“处理中”状态节点，Unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	PlanProcessingEndDate *string `json:"plan_processing_end_date,omitempty"`

	// **参数解释：**  计划研发结束时间。通常用于“研发”状态节点，Unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	PlanResearchanddevelopEndDate *string `json:"plan_researchanddevelop_end_date,omitempty"`

	// **参数解释：**  计划测试结束时间。通常用于“测试”状态节点，Unix时间戳，单位为毫秒。 **取值范围：**  不涉及。
	PlanTestEndDate *string `json:"plan_test_end_date,omitempty"`

	// **参数解释：**  标识工作项在列表中初始排序位置。 **取值范围：**  不涉及。
	PositionFloat *string `json:"position_float,omitempty"`

	// **参数解释：**  处理中责任人。通常用于“处理中”状态节点。 **取值范围：**  不涉及。
	ProcessingOwner *string `json:"processing_owner,omitempty"`

	// **参数解释：**  分析原因。 **取值范围：**  不涉及。
	ReasonAnalysis *string `json:"reason_analysis,omitempty"`

	// **参数解释：**  回归不通过次数。缺陷测试不通过时自动赋值。 **取值范围：**  不涉及。
	RegressionFailureNumber *int32 `json:"regression_failure_number,omitempty"`

	RelatedNetworkSecurity *OptionVo `json:"related_network_security,omitempty"`

	// **参数解释：**  修复方案。常用于修复缺陷时。 **取值范围：**  不涉及。
	RepairSolution *string `json:"repair_solution,omitempty"`

	// **参数解释：**  研发责任人。通常用于“研发”状态节点。 **取值范围：**  不涉及。
	ResearchanddevelopOwner *string `json:"researchanddevelop_owner,omitempty"`

	Severity *OptionVo `json:"severity,omitempty"`

	// **参数解释：**  严重程度。 **取值范围：**  不涉及。
	SysActivationReason *string `json:"sys_activation_reason,omitempty"`

	// **参数解释：**  无需修复原因。通常用于在缺陷无需修复时。 **取值范围：**  不涉及。
	SysNoRepairReason *string `json:"sys_no_repair_reason,omitempty"`

	// **参数解释：**  测试不通过次数。 **取值范围：**  不涉及。
	TestFailuresTimes *int32 `json:"test_failures_times,omitempty"`

	// **参数解释：**  测试责任人。通常用于“测试”状态节点。 **取值范围：**  不涉及。
	TestOwner *string `json:"test_owner,omitempty"`

	// **参数解释：**  测试报告。 **取值范围：**  不涉及。
	TestReport *string `json:"test_report,omitempty"`

	ValFeature *OptionVo `json:"val_feature,omitempty"`

	// **参数解释：**  工作项关联的甘特图Id。 **取值范围：**  不涉及。
	Workitem2ganttchart *string `json:"workitem2ganttchart,omitempty"`

	// **参数解释：**  工作项关联的思维导图Id。 **取值范围：**  不涉及。
	Workitem2mindmap *string `json:"workitem2mindmap,omitempty"`
}

func (o IssueVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueVo struct{}"
	}

	return strings.Join([]string{"IssueVo", string(data)}, " ")
}
