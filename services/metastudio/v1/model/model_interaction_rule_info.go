package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionRuleInfo 互动规则配置信息
type InteractionRuleInfo struct {

	// 规则索引
	RuleIndex *string `json:"rule_index,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 事件类型。 * 1：弹幕事件 * 2：用户入场事件 * 3：用户点赞事件 * 4：用户送礼事件 * 10: 预置话术事件
	EventType *int32 `json:"event_type,omitempty"`

	HitCondition *HitCondition `json:"hit_condition,omitempty"`

	Trigger *TriggerProcess `json:"trigger,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`
}

func (o InteractionRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleInfo struct{}"
	}

	return strings.Join([]string{"InteractionRuleInfo", string(data)}, " ")
}
