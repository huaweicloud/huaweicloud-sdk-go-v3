package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InteractionRuleDetailInfo struct {

	// **参数解释**： 规则索引。用于触发规则时索引具体规则。 **约束限制**： 无需用户填写。 **取值范围**： 字符0-64位 **默认取值**： 不涉及。
	RuleIndex *string `json:"rule_index,omitempty"`

	// **参数解释**： 规则名称。 **约束限制**： 不涉及。 **取值范围**： 字符0-256位 **默认取值**： 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**： 是否启用。 **约束限制**： 不涉及。 **取值范围**： * true：启用 * fasle：不启用  **默认取值**： true
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 规则匹配直播事件类型。接口的取值范围[0,100]，实际业务取值如下所示： * 1：弹幕事件  * 2：用户入场事件  * 3：用户点赞事件 * 4：用户送礼事件  * 10: 预置话术事件  请以实际业务取值为准。 > * 1,2,3,4：与LiveEventReport中的event.type对应。 > * 10：匹配预置剧本  **约束限制**： 不涉及。 **默认取值**： 不涉及
	EventType *int32 `json:"event_type,omitempty"`

	HitCondition *HitCondition `json:"hit_condition,omitempty"`

	Trigger *TriggerProcess `json:"trigger,omitempty"`

	ReviewConfig *ReviewConfig `json:"review_config,omitempty"`

	// 互动规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o InteractionRuleDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleDetailInfo struct{}"
	}

	return strings.Join([]string{"InteractionRuleDetailInfo", string(data)}, " ")
}
