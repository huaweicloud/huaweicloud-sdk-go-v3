package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMRuleStatusRequestBody **参数解释：** 更新BotM规则启用状态的请求体 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateBotMRuleStatusRequestBody struct {

	// **参数解释：** 是否开启 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Enable *bool `json:"enable,omitempty"`
}

func (o UpdateBotMRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleStatusRequestBody", string(data)}, " ")
}
