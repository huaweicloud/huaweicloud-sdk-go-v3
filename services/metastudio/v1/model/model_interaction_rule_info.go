package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionRuleInfo 互动规则配置信息
type InteractionRuleInfo struct {

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 事件类型
	EventType *int32 `json:"event_type,omitempty"`

	HitCondition *HitCondition `json:"hit_condition,omitempty"`

	Trigger *TriggerProcess `json:"trigger,omitempty"`
}

func (o InteractionRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleInfo struct{}"
	}

	return strings.Join([]string{"InteractionRuleInfo", string(data)}, " ")
}
