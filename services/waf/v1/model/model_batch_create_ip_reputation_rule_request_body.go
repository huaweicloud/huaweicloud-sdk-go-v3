package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateIpReputationRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 规则名称只能由字母、数字、-、_和.组成，长度不能超过64个字符 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 信誉类型（目前只支持idc） **约束限制：** 不涉及 **取值范围：** - idc   **默认取值：** 不涉及
	Type string `json:"type"`

	// **参数解释：** 标签列表，用于指定关联的情报标识 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags []string `json:"tags"`

	Action *UpdateIpReputationRuleRequestBodyAction `json:"action"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateIpReputationRuleRequestBody", string(data)}, " ")
}
