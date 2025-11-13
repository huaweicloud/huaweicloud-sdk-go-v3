package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInteractionRuleReq 单条互动规则增加请求
type CreateInteractionRuleReq struct {

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 事件类型。 * 1：弹幕事件 * 2：用户入场事件 * 3：用户点赞事件 * 4：用户送礼事件 * 10: 预置话术事件
	EventType *int32 `json:"event_type,omitempty"`

	HitCondition *HitCondition `json:"hit_condition,omitempty"`

	Trigger *TriggerProcess `json:"trigger,omitempty"`
}

func (o CreateInteractionRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInteractionRuleReq struct{}"
	}

	return strings.Join([]string{"CreateInteractionRuleReq", string(data)}, " ")
}
