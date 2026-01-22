package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleHitStatusListVo struct {

	// **参数解释**： 规则ID **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 规则命中次数 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RuleHitCount *int64 `json:"rule_hit_count,omitempty"`

	// **参数解释**： 命中时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	RuleLastHitTime *int64 `json:"rule_last_hit_time,omitempty"`
}

func (o RuleHitStatusListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleHitStatusListVo struct{}"
	}

	return strings.Join([]string{"RuleHitStatusListVo", string(data)}, " ")
}
