package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpReputationRuleRequestBodyAction **参数解释：** 防护动作配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateIpReputationRuleRequestBodyAction struct {

	// **参数解释：** 动作类型（如captcha表示验证码） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Category *string `json:"category,omitempty"`
}

func (o UpdateIpReputationRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpReputationRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"UpdateIpReputationRuleRequestBodyAction", string(data)}, " ")
}
