package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InteractionRuleDetailInfo struct {

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

	// 互动规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o InteractionRuleDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleDetailInfo struct{}"
	}

	return strings.Join([]string{"InteractionRuleDetailInfo", string(data)}, " ")
}
