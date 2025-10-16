package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateIpReputationRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 策略名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policyname *string `json:"policyname,omitempty"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	Action *UpdateIpReputationRuleRequestBodyAction `json:"action,omitempty"`

	// **参数解释：** 规则类型（如idc表示机房IP情报类型） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释：** 标签列表，关联的情报标识 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags *[]string `json:"tags,omitempty"`

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds *[]PolicyRuleIdRequestBodyPolicyRuleIds `json:"policy_rule_ids,omitempty"`
}

func (o BatchUpdateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateIpReputationRuleRequestBody", string(data)}, " ")
}
