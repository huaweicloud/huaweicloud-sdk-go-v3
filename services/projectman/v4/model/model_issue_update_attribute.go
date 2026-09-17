package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueUpdateAttribute 工作项更新对象
type IssueUpdateAttribute struct {

	// **参数解释**： 工作项类型编码。编辑工作项时，此字段必填、值为当前工作项正确的工作项类型，但不会更新此字段。 **约束限制**： 不涉及。 **取值范围**： 支持多种工作项类型，使用英文逗号分隔。 - 系统设备类项目：RR、SF、IR、SR、AR、Task、Bug - 独立软件类项目：RR、SF、IR、US、Task、Bug - 云服务类项目：RR、Epic、FE、US、Task、Bug **默认取值**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 工作项描述字段，可通过[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**description**字段的值就是工作项描述字段。 **约束限制**： 不涉及。 **取值范围**： 1~500000个字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 父工作项ID，可通过[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**parent_id**字段的值就是父工作项ID。 **约束限制**： 不涉及。 **取值范围**： 18~19个字符的数字字符串。 **默认取值**： 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 工作项状态code。可通过[查询工作项状态](ListIssueStatues.xml)接口获取，响应消息体中的**code**字段的值就是工作项状态code。 **约束限制**： 不涉及。 **取值范围**： 2~32个字符。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	Assignee *UserUpdateAttribute `json:"assignee,omitempty"`

	// **参数解释**： 工作项抄送人，支持多个抄送人。数组元素为UserUpdateAttribute对象。 **约束限制**： 同一工作项最多支持50个抄送人。
	AssignedCc *[]UserUpdateAttribute `json:"assigned_cc,omitempty"`

	// **参数解释**： 工作项计划结束日期，unix时间戳，单位：毫秒。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// **参数解释**： 工作项计划工时。 **约束限制**： 保留一位小数。 **取值范围**： 0~999999999.9。 **默认取值**： 不涉及。
	Workload *string `json:"workload,omitempty"`

	// **参数解释**： 工作项关联项ID。 **约束限制**： 多个关联项用英文逗号分隔，同一工作项最多支持50个关联项。 **取值范围**： 0~2048个字符。 **默认取值**： 不涉及。
	Link *string `json:"link,omitempty"`

	// **参数解释**： 工作项标签。数组元素为LabelEntity对象。 **约束限制**： 不涉及。 **取值范围**： 0~50个元素，每个元素为LabelEntity对象。 **默认取值**： 不涉及。
	Labels *[]LabelEntity `json:"labels,omitempty"`

	// **参数解释**： 工作项自定义字段映射，用户添加的系统字段也在此列，格式为{\"code\":\"字段code\",\"value\":\"字段值\"}。数组元素为FieldCodeValuePair对象。 **约束限制**： 不涉及。 **取值范围**： 0~200个元素，每个元素为FieldCodeValuePair对象。 **默认取值**： 不涉及。
	CustomFields *[]FieldCodeValuePair `json:"custom_fields,omitempty"`

	// **参数解释**： IR和SF的关联字段。 **约束限制**： IR可以填写该字段。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ir2feature *string `json:"ir2feature,omitempty"`

	// **参数解释**： 工作项是否需要分解。 **约束限制**： 仅可以分解的工作项类型有此字段。 **取值范围**： - yes：需要分解 - no：不需要分解 **默认取值**： 不涉及。
	NeedBreak *string `json:"need_break,omitempty"`

	// **参数解释**： 工作项基线状态。 **约束限制**： 不涉及。 **取值范围**： - null：未基线 - baselined：已基线 - baseline-reviewing：基线评审中 **默认取值**： 不涉及。
	Baseline *string `json:"baseline,omitempty"`

	// **参数解释**： 工作项优先级，部分工作项有此字段。 **约束限制**： 不涉及。 **取值范围**： - 低：低优先级。 - 中：中优先级。 - 高：高优先级。 **默认取值**： 不涉及。
	Priority *string `json:"priority,omitempty"`

	// **参数解释**： 是否涉及网络安全。 **约束限制**： 预设字段中，仅研发需求类型的工作项有此字段。 **取值范围**： - yes：涉及网络安全。 - no：不涉及网络安全。 **默认取值**： 不涉及。
	RelatedNetworkSecurity *string `json:"related_network_security,omitempty"`

	// **参数解释**： 领域字段。 **约束限制**： 不涉及。 **取值范围**： - software - soft-hardware - hardware - 性能 - 功能 - 运维 - 运营 - 用户体验 - 隐私保护 - 合规 - 韧性(可靠性/可用性) - 韧性(危险检测与相应恢复) - 透明 - 无害 - 安全 - API - 成本 - 可维护性 - 其他DFX - 可用性 - others **默认取值**： 不涉及。
	BusinessDomain *string `json:"business_domain,omitempty"`

	// **参数解释**： 工作项发布计划ID。通过[发布/迭代计划列表查询](ListPlan.xml)接口查询计划列表，返回参数中PlanVO里面的category=PI的对象的**id**字段就是迭代计划的ID。 **约束限制**： 不涉及。 **取值范围**： 18~19个字符的数字字符串。 **默认取值**： 不涉及。
	PlanPi *string `json:"plan_pi,omitempty"`

	// **参数解释**： 工作项完成的迭代计划ID，在Bug中为修复迭代计划ID。通过[发布/迭代计划列表查询](ListPlan.xml)接口查询计划列表，返回参数中PlanVO里面的category=Iteration的对象的**id**字段就是迭代计划的ID。 **约束限制**： 18~19个字符的数字字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanIteration *string `json:"plan_iteration,omitempty"`

	// **参数解释**： 无需分解原因。 **约束限制**： need_break字段值为“no”时有此字段。 **取值范围**： 0~512个字符。 **默认取值**： 不涉及。
	NoBreakReason *string `json:"no_break_reason,omitempty"`

	// **参数解释**： 工作项提出人。数组元素为UserUpdateAttribute对象。 **约束限制**： 不涉及。
	SubmittedBy *[]UserUpdateAttribute `json:"submitted_by,omitempty"`

	// **参数解释**： IR关联的RR ID，可以通过[查询工作项列表](ListIpdProjectIssues.xml)或者[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**id**字段的值就是工作项ID。 **约束限制**： 多个关联项ID使用英文逗号分隔。 **取值范围**： 0~1024个字符。 **默认取值**： 不涉及。
	Ir2rr *string `json:"ir2rr,omitempty"`

	// **参数解释**： 特性集ID，可以通过[查询特性集](ShowBaselineSnapshots.xml)接口获取，响应消息体中的**id**字段的值就是特性集ID。 **约束限制**： 不涉及。 **取值范围**： 18~19个字符的数字字符串。 **默认取值**： 不涉及。
	FeatureSet *string `json:"feature_set,omitempty"`

	// **参数解释**： 期望修复时间。预设字段中，仅Bug有此字段，unix时间戳，单位：毫秒。 **约束限制**： 不涉及。 **取值范围**： 11~19个字符。 **默认取值**： 不涉及。
	ExpectedRepairDate *string `json:"expected_repair_date,omitempty"`

	// **参数解释**： 缺陷发现发布计划ID，预设字段中，仅Bug有此字段。通过[发布/迭代计划列表查询](ListPlan.xml)接口查询计划列表，返回参数中PlanVO里面的category=PI的对象的**id**字段就是迭代计划的ID。 **约束限制**： 不涉及。 **取值范围**： 18~19个字符的数字字符串。 **默认取值**： 不涉及。
	FoundPi *string `json:"found_pi,omitempty"`

	// **参数解释**： 缺陷发现迭代计划ID，预设字段中，仅Bug有此字段。通过[发布/迭代计划列表查询](ListPlan.xml)接口查询计划列表，返回参数中PlanVO里面的category=Iteration的对象的**id**字段就是迭代计划的ID。 **约束限制**： 不涉及。 **取值范围**： 18~19个字符的数字字符串。 **默认取值**： 不涉及。
	FoundIteration *string `json:"found_iteration,omitempty"`

	// **参数解释**： 分析原因。 **约束限制**： 预设字段中，仅Bug有此字段。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	ReasonAnalysis *string `json:"reason_analysis,omitempty"`

	// **参数解释**： 修复方案。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	RepairSolution *string `json:"repair_solution,omitempty"`

	// **参数解释**： 测试报告。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	TestReport *string `json:"test_report,omitempty"`

	// **参数解释**： 无需修复原因。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	SysNoRepairReason *string `json:"sys_no_repair_reason,omitempty"`

	// **参数解释**： 激活原因。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	SysActivationReason *string `json:"sys_activation_reason,omitempty"`

	// **参数解释**： 退回原因。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	SysReturnReason *string `json:"sys_return_reason,omitempty"`

	// **参数解释**： 回归不通过次数。预设字段中，仅Bug有此字段。 **约束限制**： 不涉及。 **取值范围**： 0~999999。 **默认取值**： 不涉及。
	TestFailuresTimes *int32 `json:"test_failures_times,omitempty"`

	// **参数解释**： 关闭类型。 **约束限制**： 不涉及。 **取值范围**： - problem_solved：问题解决关闭 - problem_to_requirement：问题转需求关闭 - duplicate_problem：重复问题关闭 - not_a_problem：非问题关闭 **默认取值**： 不涉及。
	CloseType *string `json:"close_type,omitempty"`

	// **参数解释**： 密级。低密级权限者不能访问高密级的工作项。可以通过[查询字段列表](ListIpdProjectFields.xml)接口获取，响应消息体中密级的**option**字段的值就是密级字段的可选值。 **约束限制**： 仅在涉密环境（SM）下存在此字段，非涉密环境下无此字段。涉密环境下必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityLevel *string `json:"security_level,omitempty"`

	PlanOwner *UserUpdateAttribute `json:"plan_owner,omitempty"`

	DoingOwner *UserUpdateAttribute `json:"doing_owner,omitempty"`

	DeliveredOwner *UserUpdateAttribute `json:"delivered_owner,omitempty"`

	CheckingOwner *UserUpdateAttribute `json:"checking_owner,omitempty"`

	TestOwner *UserUpdateAttribute `json:"test_owner,omitempty"`

	DevelopOwner *UserUpdateAttribute `json:"develop_owner,omitempty"`

	ProcessingOwner *UserUpdateAttribute `json:"processing_owner,omitempty"`

	FixedOwner *UserUpdateAttribute `json:"fixed_owner,omitempty"`

	ResearchanddevelopOwner *UserUpdateAttribute `json:"researchanddevelop_owner,omitempty"`

	AnalyseOwner *UserUpdateAttribute `json:"analyse_owner,omitempty"`

	// **参数解释**： 计划开始时间。工作项的计划启动日期，用于项目进度管理和排期。 **约束限制**： 不涉及。 **取值范围**： 11~19个字符的时间戳字符串，单位为毫秒（ms）。 **默认取值**： 不涉及。
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// **参数解释**： 期望完成时间。工作项的预期交付日期，用于跟踪工作项是否按期完成。 **约束限制**： 不涉及。 **取值范围**： 11~19个字符的时间戳字符串，单位为毫秒（ms）。 **默认取值**： 不涉及。
	ExpectDeliveryTime *string `json:"expect_delivery_time,omitempty"`

	// **参数解释**： 计划测试结束时间。Bug类型工作项的计划测试完成日期，用于跟踪Bug修复后的测试进度。 **约束限制**： 仅对Bug类型工作项生效，非Bug类型忽略此字段。 **取值范围**： 11~19个字符的时间戳字符串，单位为毫秒（ms）。 **默认取值**： 不涉及。
	PlanTestEndDate *string `json:"plan_test_end_date,omitempty"`

	// **参数解释**： 严重程度。Bug类型工作项的严重级别，用于评估Bug的影响范围和修复优先级。 **约束限制**： 仅对Bug类型工作项生效，非Bug类型忽略此字段。 **取值范围**： - 致命：系统崩溃、数据丢失等严重影响 - 严重：主要功能无法使用 - 一般：次要功能异常，有替代方案 - 提示：界面优化、建议性问题 **默认取值**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 是否承诺。RR（原始需求）类型工作项的承诺状态标识，用于标记需求是否已承诺交付。 **约束限制**： 仅对RR类型工作项生效，非RR类型忽略此字段。 **取值范围**： - yes：已承诺 - no：未承诺 **默认取值**： 不涉及。
	Promised *string `json:"promised,omitempty"`

	// **参数解释**： 承接人。RR（原始需求）类型工作项的需求承接责任人，负责需求的分析和转化。 **约束限制**： 仅对RR类型工作项生效，非RR类型忽略此字段。
	Recipient *[]UserUpdateAttribute `json:"recipient,omitempty"`

	// **参数解释**： 无需研发原因。RR（原始需求）类型工作项不需要进行研发的原因说明。 **约束限制**： 仅对RR类型工作项生效，非RR类型忽略此字段。 **取值范围**： 0~50000个字符。 **默认取值**： 不涉及。
	SysNoDevelopReason *string `json:"sys_no_develop_reason,omitempty"`

	// **参数解释**： 价值特性。SF/FE类型工作项对应的业务价值特性描述，用于关联业务价值和技术实现。 **约束限制**： 仅对SF/FE类型工作项生效，其他类型忽略此字段。 **取值范围**： - yes：是 - no：否 **默认取值**： 不涉及。
	ValFeature *string `json:"val_feature,omitempty"`

	// **参数解释**： 功能场景。SF/FE类型工作项的功能应用场景描述，用于说明特性的使用场景和用户故事。 **约束限制**： 仅对SF/FE类型工作项生效，其他类型忽略此字段。 **取值范围**： 0~512个字符。 **默认取值**： 不涉及。
	FunctionScene *string `json:"function_scene,omitempty"`
}

func (o IssueUpdateAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueUpdateAttribute struct{}"
	}

	return strings.Join([]string{"IssueUpdateAttribute", string(data)}, " ")
}
