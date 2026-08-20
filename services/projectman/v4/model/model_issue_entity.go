package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueEntity 工作项对象
type IssueEntity struct {

	// 需要更新的工作项ID，可通过查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	Id *string `json:"id,omitempty"`

	// 工作项标题，可通过查询树状工作项接口获取，响应消息体中的title字段的值就是工作项标题。
	Title *string `json:"title,omitempty"`

	// 工作项描述字段，可通过查询树状工作项接口获取，响应消息体中的description字段的值就是工作项描述字段。
	Description *string `json:"description,omitempty"`

	// 工作项大分类定义。工作项创建、编辑无此字段，仅作展示用，可通过查询树状工作项接口获取，响应消息体中的type字段的值就是工作项大分类定义。
	Type *string `json:"type,omitempty"`

	// 工作项编号，可通过查询树状工作项接口获取，响应消息体中的number字段的值就是工作项编号。
	Number *string `json:"number,omitempty"`

	// 工作项类型，可通过查询树状工作项接口获取，响应消息体中的category字段的值就是工作项类型。
	Category string `json:"category"`

	// 工作项类型层级关系ID，此参数影响工作项的层级显示。通过获取模型树配置信息获取，根据参数中的category在响应消息体中category_layer_config中找到对应的category_code，和category_code同级的id就是工作项类型层级关系ID。
	CategoryLayerId string `json:"category_layer_id"`

	// 父工作项ID，可通过查询树状工作项接口获取，响应消息体中的parent_id字段的值就是父工作项ID。
	ParentId *string `json:"parent_id,omitempty"`

	// 项目的32位uuid，项目唯一标识，通过查询IPD项目列表接口获取，响应消息体中的project_id字段的值就是项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 工作项状态code。可通过查询工作项状态接口获取，响应消息体中的code字段的值就是工作项工作项状态code。
	Status *string `json:"status,omitempty"`

	// 工作项的生命周期，可选值为“正在工作”，“作废”，可通过查询树状工作项接口获取，响应消息体中的state字段的值就是工作项的生命周期。
	State *string `json:"state,omitempty"`

	Assignee *UserEntity `json:"assignee,omitempty"`

	// 工作项抄送人，支持多个抄送人。数组元素为UserEntity对象。
	AssignedCc *[]UserEntity `json:"assigned_cc,omitempty"`

	CreatedBy *UserEntity `json:"created_by,omitempty"`

	// 工作项创建时间，unix时间戳，单位：毫秒。
	CreatedTime *string `json:"created_time,omitempty"`

	ModifiedBy *UserEntity `json:"modified_by,omitempty"`

	// 工作项最近更新时间，unix时间戳，单位：毫秒。
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 工作项计划结束日期，unix时间戳，单位：毫秒。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// 工作项关闭时间，unix时间戳，单位：毫秒。
	CloseTime *string `json:"close_time,omitempty"`

	// 工作项计划工时。
	Workload *string `json:"workload,omitempty"`

	// 工作项实际工时。
	WorkloadSum *string `json:"workload_sum,omitempty"`

	// 工作项所属租户ID，可通过查询树状工作项接口获取，响应消息体中的tenant_id字段的值就是工作项所属租户ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 工作项关联项ID。
	Link *string `json:"link,omitempty"`

	// 工作项是否已挂起。
	Suspended *bool `json:"suspended,omitempty"`

	// 工作项状态改变时间，可用于计算工作项在当前状态停留天数，unix时间戳，单位：毫秒。
	StatusModifiedTime *string `json:"status_modified_time,omitempty"`

	// 工作项标签。数组元素为LabelEntity对象。
	Labels *[]LabelEntity `json:"labels,omitempty"`

	// 工作项自定义字段映射，用户添加的系统字段也在此列，格式为{\"code\":\"字段code\",\"value\":\"字段值\"}。数组元素为FieldCodeValuePair对象。
	CustomFields *[]FieldCodeValuePair `json:"custom_fields,omitempty"`

	// 工作项的子工作项集合。数组元素为IssueEntity对象。
	Children *[]IssueEntity `json:"children,omitempty"`

	// 子工作项的路径。
	Path *string `json:"path,omitempty"`

	// IR和SF的关联字段。
	Ir2feature *string `json:"ir2feature,omitempty"`

	// 工作项是否需要分解。
	NeedBreak *string `json:"need_break,omitempty"`

	// 分解状态。
	BreakStatus *string `json:"break_status,omitempty"`

	// 工作项基线状态。
	Baseline *string `json:"baseline,omitempty"`

	// 工作项优先级，部分工作项有此字段。
	Priority *string `json:"priority,omitempty"`

	// 是否涉及网络安全。
	RelatedNetworkSecurity *string `json:"related_network_security,omitempty"`

	// 研发需求协同信息，协同任务ID，可通过查询树状工作项接口获取，响应消息体中的collaboratives字段的值就是研发需求协同信息，协同任务ID。
	Collaboratives *string `json:"collaboratives,omitempty"`

	// 领域字段。
	BusinessDomain *string `json:"business_domain,omitempty"`

	// 工作项发布计划ID。通过发布/迭代计划列表查询接口查询计划列表，返回参数中PlanVO里面的category=PI的对象的id字段就是迭代计划的ID。
	PlanPi *string `json:"plan_pi,omitempty"`

	// 工作项完成的迭代计划ID，在Bug中为修复迭代计划ID。通过发布/迭代计划列表查询接口查询计划列表，返回参数中PlanVO里面的category=Iteration的对象的id字段就是迭代计划的ID。
	PlanIteration *string `json:"plan_iteration,omitempty"`

	// 工作项变更状态。
	ChangeStatus *string `json:"change_status,omitempty"`

	// 无需分解原因。
	NoBreakReason *string `json:"no_break_reason,omitempty"`

	// 工作项提出人。数组元素为UserEntity对象。
	SubmittedBy *[]UserEntity `json:"submitted_by,omitempty"`

	// IR关联的RR ID，可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	Ir2rr *string `json:"ir2rr,omitempty"`

	// 特性集ID，可以通过查询特性集接口获取，响应消息体中的id字段的值就是特性集ID。
	FeatureSet *string `json:"feature_set,omitempty"`

	// 期望修复时间。预设字段中，仅Bug有此字段，unix时间戳，单位：毫秒。
	ExpectedRepairDate *string `json:"expected_repair_date,omitempty"`

	// 缺陷发现发布计划ID，预设字段中，仅Bug有此字段。通过发布/迭代计划列表查询接口查询计划列表，返回参数中PlanVO里面的category=PI的对象的id字段就是迭代计划的ID。
	FoundPi *string `json:"found_pi,omitempty"`

	// 缺陷发现迭代计划ID，预设字段中，仅Bug有此字段。通过发布/迭代计划列表查询接口查询计划列表，返回参数中PlanVO里面的category=Iteration的对象的id字段就是迭代计划的ID。
	FoundIteration *string `json:"found_iteration,omitempty"`

	// 分析原因。
	ReasonAnalysis *string `json:"reason_analysis,omitempty"`

	// 修复方案。预设字段中，仅Bug有此字段。
	RepairSolution *string `json:"repair_solution,omitempty"`

	// 测试报告。预设字段中，仅Bug有此字段。
	TestReport *string `json:"test_report,omitempty"`

	// 无需修复原因。预设字段中，仅Bug有此字段。
	SysNoRepairReason *string `json:"sys_no_repair_reason,omitempty"`

	// 激活原因。预设字段中，仅Bug有此字段。
	SysActivationReason *string `json:"sys_activation_reason,omitempty"`

	// 退回原因。预设字段中，仅Bug有此字段。
	SysReturnReason *string `json:"sys_return_reason,omitempty"`

	// 回归不通过次数。预设字段中，仅Bug有此字段。
	TestFailuresTimes *int32 `json:"test_failures_times,omitempty"`

	// 关闭类型。
	CloseType *string `json:"close_type,omitempty"`

	PlanOwner *UserEntity `json:"plan_owner,omitempty"`

	DoingOwner *UserEntity `json:"doing_owner,omitempty"`

	DeliveredOwner *UserEntity `json:"delivered_owner,omitempty"`

	CheckingOwner *UserEntity `json:"checking_owner,omitempty"`

	TestOwner *UserEntity `json:"test_owner,omitempty"`

	DevelopOwner *UserEntity `json:"develop_owner,omitempty"`

	ProcessingOwner *UserEntity `json:"processing_owner,omitempty"`

	FixedOwner *UserEntity `json:"fixed_owner,omitempty"`

	ResearchanddevelopOwner *UserEntity `json:"researchanddevelop_owner,omitempty"`

	AnalyseOwner *UserEntity `json:"analyse_owner,omitempty"`

	// 计划开始时间。工作项的计划启动日期，用于项目进度管理和排期。
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// 期望完成时间。工作项的预期交付日期，用于跟踪工作项是否按期完成。
	ExpectDeliveryTime *string `json:"expect_delivery_time,omitempty"`

	// 计划测试结束时间。Bug类型工作项的计划测试完成日期，用于跟踪Bug修复后的测试进度。
	PlanTestEndDate *string `json:"plan_test_end_date,omitempty"`

	// 严重程度。Bug类型工作项的严重级别，用于评估Bug的影响范围和修复优先级。
	Severity *string `json:"severity,omitempty"`

	// 是否承诺。RR（原始需求）类型工作项的承诺状态标识，用于标记需求是否已承诺交付。
	Promised *string `json:"promised,omitempty"`

	// 承接人。RR（原始需求）类型工作项的需求承接责任人，负责需求的分析和转化。
	Recipient *[]UserEntity `json:"recipient,omitempty"`

	// 无需研发原因。RR（原始需求）类型工作项不需要进行研发的原因说明。
	SysNoDevelopReason *string `json:"sys_no_develop_reason,omitempty"`

	// 价值特性。SF/FE类型工作项对应的业务价值特性描述，用于关联业务价值和技术实现。
	ValFeature *string `json:"val_feature,omitempty"`

	// 功能场景。SF/FE类型工作项的功能应用场景描述，用于说明特性的使用场景和用户故事。
	FunctionScene *string `json:"function_scene,omitempty"`
}

func (o IssueEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueEntity struct{}"
	}

	return strings.Join([]string{"IssueEntity", string(data)}, " ")
}
