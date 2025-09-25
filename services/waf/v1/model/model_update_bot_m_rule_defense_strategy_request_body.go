package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMRuleDefenseStrategyRequestBody **参数解释：** 更新BotM防护策略的请求体，包含低、中、高三个风险等级的防护配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdateBotMRuleDefenseStrategyRequestBody struct {
	Low *BotMDefenseLevel `json:"low,omitempty"`

	Medium *BotMDefenseLevel `json:"medium,omitempty"`

	High *BotMDefenseLevel `json:"high,omitempty"`
}

func (o UpdateBotMRuleDefenseStrategyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleDefenseStrategyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleDefenseStrategyRequestBody", string(data)}, " ")
}
