package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveRoomInteractionRuleInfo 直播使用互动规则配置信息
type LiveRoomInteractionRuleInfo struct {

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

	// 互动规则库ID。从互动库导入时携带互动库ID。
	GroupId *string `json:"group_id,omitempty"`

	// 互动规则库名称。从互动库导入时携带互动库名称。
	GroupName *string `json:"group_name,omitempty"`

	// 规则启用场景。部分场景启用时填写。不填写为全部。
	EnabledScenes *[]string `json:"enabled_scenes,omitempty"`
}

func (o LiveRoomInteractionRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveRoomInteractionRuleInfo struct{}"
	}

	return strings.Join([]string{"LiveRoomInteractionRuleInfo", string(data)}, " ")
}
