package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowsInfoVo 工作项流转线信息。
type FlowsInfoVo struct {

	// **参数解释**： 流转线code。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 流转线名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 流转线描述信息。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 流转线扩展配置。 **取值范围**： 不涉及。
	ExtraConfig *[]map[string]interface{} `json:"extra_config,omitempty"`

	// **参数解释**： 当前工作流节点code。 **取值范围**： 不涉及。
	FromCode *string `json:"from_code,omitempty"`

	// **参数解释**： 目标工作流节点code。 **取值范围**： 不涉及。
	ToCode *string `json:"to_code,omitempty"`

	// **参数解释**： 流转前规则配置。 **取值范围**： 不涉及。
	BeforeRuleConfigs *[]WorkItemFlowRuleConfigVo `json:"before_rule_configs,omitempty"`

	// **参数解释**： 流转前校验规则。 **取值范围**： 不涉及。
	BeforeRuleValidator *[]string `json:"before_rule_validator,omitempty"`

	// **参数解释**： 流转后规则配置。 **取值范围**： 不涉及。
	AfterRuleConfigs *[]WorkItemFlowRuleConfigVo `json:"after_rule_configs,omitempty"`
}

func (o FlowsInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowsInfoVo struct{}"
	}

	return strings.Join([]string{"FlowsInfoVo", string(data)}, " ")
}
