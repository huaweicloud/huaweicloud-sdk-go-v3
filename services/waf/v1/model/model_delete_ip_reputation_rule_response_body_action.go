package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpReputationRuleResponseBodyAction **参数解释：** 规则删除前的防护动作配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type DeleteIpReputationRuleResponseBodyAction struct {

	// **参数解释：** 动作类型（如log表示仅记录） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Category *string `json:"category,omitempty"`
}

func (o DeleteIpReputationRuleResponseBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpReputationRuleResponseBodyAction struct{}"
	}

	return strings.Join([]string{"DeleteIpReputationRuleResponseBodyAction", string(data)}, " ")
}
