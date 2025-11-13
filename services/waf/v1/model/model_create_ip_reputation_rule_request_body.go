package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateIpReputationRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 规则名称只能由字母、数字、-、_和.组成，长度不能超过64个字符 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 策略名称 **约束限制：** 不涉及 **取值范围：** 策略名称只能由字母、数字、-、_和.组成，长度不能超过64个字符 **默认取值：** 不涉及
	Policyname *string `json:"policyname,omitempty"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态（1表示开启，0表示关闭） **约束限制：** 不涉及 **取值范围：** - 0：关闭 - 1：开启  **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	Action *UpdateIpReputationRuleRequestBodyAction `json:"action"`

	// **参数解释：** 信誉类型（目前只支持idc） **约束限制：** 不涉及 **取值范围：** - idc   **默认取值：** 不涉及
	Type string `json:"type"`

	// **参数解释：** 标签列表，用于指定关联的情报标识 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags []string `json:"tags"`
}

func (o CreateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpReputationRuleRequestBody", string(data)}, " ")
}
