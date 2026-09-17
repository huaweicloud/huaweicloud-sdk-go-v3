package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineDto struct {

	// **参数解释**： 流水线名称。 **约束限制**： 不涉及。 **取值范围**： 仅包含中文、大小写英文字母、数字、'-'和'_'，且长度为[1,128]个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 流水线描述。 **约束限制**： 不涉及。 **取值范围**： 不超过1024字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否为变更流水线。 **约束限制**： 不涉及。 **取值范围**： - true：变更流水线。 - false：非变更流水线。 **默认取值**： 不涉及。
	IsPublish bool `json:"is_publish"`

	// **参数解释**： 流水线源信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sources *[]CodeSource `json:"sources,omitempty"`

	// **参数解释**： 流水线自定义全局变量列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Variables *[]CustomVariable `json:"variables,omitempty"`

	// **参数解释**： 流水线定时执行配置列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Schedules *[]PipelineSchedule `json:"schedules,omitempty"`

	// **参数解释**： 流水线代码事件触发配置。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Triggers *[]PipelineTrigger `json:"triggers,omitempty"`

	// **参数解释**： 流水线结构定义版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 3.0。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 流水线结构定义JSON。该字段结构复杂，建议使用页面编辑流水线后，从[查询流水线详情](ShowPipelineDetail.xml)接口获取。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Definition string `json:"definition"`

	// **参数解释**： 项目名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**： 流水线组ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 复制场景使用，为原流水线ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	ConcurrencyControl *PipelineConcurrencyMgmt `json:"concurrency_control,omitempty"`

	// **参数解释**： 流水线涉密等级。 **约束限制**： 非涉密场景不涉及，涉密场景必填。 **取值范围**： 正整数（1为最低密级）。 **默认取值**： 不涉及。
	SecurityLevel *int32 `json:"security_level,omitempty"`

	// **参数解释**： 是否禁用自动生成阶段，仅变更流水线适用。默认值为false，即默认启用自动生成阶段。 **约束限制**： 不涉及。 **取值范围**： - true：禁用自动生成阶段。 - false：启用自动生成阶段。 **默认取值**： false。
	DisableReleaseBranchManagement *bool `json:"disable_release_branch_management,omitempty"`

	// **参数解释**： 流水线执行计划列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ExecutionPlans *[]interface{} `json:"execution_plans,omitempty"`

	// **参数解释**： 项目ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，仅由数字和字母组成。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 流水线取消运行策略。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CancelStrategy *interface{} `json:"cancel_strategy,omitempty"`

	// **参数解释**： 流水线密级code。 **约束限制**： 非涉密场景不涉及，涉密场景必填。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConfidentialityCode *string `json:"confidentiality_code,omitempty"`

	// **参数解释**： 委托名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AgencyName *string `json:"agency_name,omitempty"`

	// **参数解释**： 关联的通用参数组ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VariableGroupIds *[]string `json:"variable_group_ids,omitempty"`
}

func (o PipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineDto struct{}"
	}

	return strings.Join([]string{"PipelineDto", string(data)}, " ")
}
