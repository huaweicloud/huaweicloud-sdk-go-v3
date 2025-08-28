package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpReputationRuleRequestBodyAction **参数解释：** 防护动作配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type CreateIpReputationRuleRequestBodyAction struct {

	// **参数解释：** 动作类型（如log表示仅记录） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Category *string `json:"category,omitempty"`
}

func (o CreateIpReputationRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpReputationRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"CreateIpReputationRuleRequestBodyAction", string(data)}, " ")
}
