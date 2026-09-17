package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueCreateEntity 工作项创建对象
type IssueCreateEntity struct {

	// **参数解释**： 工作项标题。 **约束限制**：  不涉及。 **取值范围**： 2~256个字符。 **默认取值**： 不涉及。
	Title string `json:"title"`

	// **参数解释**： 工作项描述字段。 **约束限制**： 不涉及。 **取值范围**： 0~500000个字符。 **默认取值**： 不涉及。
	Description string `json:"description"`

	// **参数解释**： 工作项类型编码。编辑工作项时，此字段必填、值为当前工作项正确的工作项类型，但不会更新此字段。 **约束限制**： 不涉及。 **取值范围**： 支持多种工作项类型，使用英文逗号分隔。 - 系统设备类项目：RR、SF、IR、SR、AR、Task、Bug - 独立软件类项目：RR、SF、IR、US、Task、Bug - 云服务类项目：RR、Epic、FE、US、Task、Bug **默认取值**： 不涉及。
	Category string `json:"category"`

	// **参数解释**： 工作项类型层级关系ID，此参数影响工作项的层级显示。通过[获取模型树配置信息](GetModelConfig.xml)获取，根据参数中的category在响应消息体中category_layer_config中找到对应的category_code，和category_code同级的id就是工作项类型层级关系ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CategoryLayerId string `json:"category_layer_id"`

	// **参数解释**： 父工作项ID。 **约束限制**： 创建子工作项时必填，其他场景非必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ParentId string `json:"parent_id"`

	// **参数解释**： 工作项状态code。可通过[查询工作项状态](ListIssueStatues.xml)接口获取，响应消息体中的**code**字段的值就是工作项状态code。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Status string `json:"status"`

	Assignee *UserEntity `json:"assignee"`

	// **参数解释**： 原始需求承接人。 **约束限制**： 当工作项类型为RR时字段必填，其他工作项类型无此字段。
	Recipient *[]UserEntity `json:"recipient,omitempty"`

	// **参数解释**： 工作项抄送人，支持多个抄送人。 **约束限制**： 同一工作项最多支持50个抄送人。
	AssignedCc *[]UserEntity `json:"assigned_cc,omitempty"`

	// **参数解释**： 工作项计划结束日期。 **约束限制**： 0~13个字符的数字字符串，可选负号前缀。 **取值范围**： 时间戳。 **默认取值**： 不涉及。
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// **参数解释**： 工作项计划工时。 **约束限制**： 不涉及。 **取值范围**： 0~999999999.9中的数字字符串。 **默认取值**： 不涉及。
	Workload *string `json:"workload,omitempty"`

	// **参数解释**： 工作项关联项ID。 **约束限制**： 多个关联项用英文逗号分隔，同一工作项最多支持50个关联项。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Link *string `json:"link,omitempty"`

	// **参数解释**： 工作项标签。 **约束限制**： 不涉及。
	Labels *[]LabelEntity `json:"labels,omitempty"`

	// **参数解释**： 工作项自定义字段映射。用户添加的系统字段也在此列。 **约束限制**： 不涉及。
	CustomFields *[]FieldCodeValuePair `json:"custom_fields,omitempty"`

	// **参数解释**： IR和SF的关联字段。 **约束限制**： IR可以填写该字段。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ir2feature *string `json:"ir2feature,omitempty"`

	// **参数解释**： 工作项优先级。 **约束限制**： RR、SF、FE、IR、SR、AR、Task、Bug可以填写该字段。 **取值范围**： - 低：低优先级。 - 中：中优先级。 - 高：高优先级。 **默认取值**： 不涉及。
	Priority *string `json:"priority,omitempty"`

	// **参数解释**： 是否涉及网络安全。 **约束限制**： 仅研发需求有此字段。 **取值范围**： - yes：涉及网络安全。 - no：不涉及网络安全。 **默认取值**： 不涉及。
	RelatedNetworkSecurity *string `json:"related_network_security,omitempty"`

	// **参数解释**： 研发需求协同信息，协同任务ID，可通过[查询树状工作项](ShowIpdIssueTree.xml)接口获取，响应消息体中的**collaboratives**字段的值就是研发需求协同信息，协同任务ID。 **约束限制**： 协同任务ID。IR、SR、AR、US有此字段。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Collaboratives *string `json:"collaboratives,omitempty"`

	// **参数解释**： 领域字段。 **约束限制**：  FE、SF、IR、SR、AR、Bug有此字段。 **取值范围**： - software - soft-hardware - hardware - 性能 - 功能 - 运维 - 运营 - 用户体验 - 隐私保护 - 合规 - 韧性(可靠性/可用性) - 韧性(危险检测与相应恢复) - 透明 - 无害 - 安全 - API - 成本 - 可维护性 - 其他DFX - 可用性 - others **默认取值**： 不涉及。
	BusinessDomain *string `json:"business_domain,omitempty"`

	// **参数解释**： 工作项发布计划ID。 **约束限制**： 默认SR、AR、US、Task、Bug有此字段。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanPi *string `json:"plan_pi,omitempty"`

	// **参数解释**： 工作项提出人。 **约束限制**： 仅RR、Bug有此字段。
	SubmittedBy *[]UserEntity `json:"submitted_by,omitempty"`

	// **参数解释**： IR关联的RR的Id。 **约束限制**： 仅IR有此字段，多选时用英文逗号分隔。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ir2rr *string `json:"ir2rr,omitempty"`

	// **参数解释**： 特性集ID。 **约束限制**： 仅SF/FE有此字段。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FeatureSet *string `json:"feature_set,omitempty"`

	// **参数解释**： 密级。低密级权限者不能访问高密级的工作项。可以通过[[查询字段列表](ListIpdProjectFields.xml)]接口获取，响应消息体中密级的**option**字段的值就是密级字段的可选值。 **约束限制**： 仅在涉密环境（SM）下存在此字段，非涉密环境下无此字段。涉密环境下必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityLevel *string `json:"security_level,omitempty"`
}

func (o IssueCreateEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCreateEntity struct{}"
	}

	return strings.Join([]string{"IssueCreateEntity", string(data)}, " ")
}
