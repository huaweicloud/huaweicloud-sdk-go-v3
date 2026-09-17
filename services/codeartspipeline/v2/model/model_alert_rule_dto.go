package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleDto 告警规则DTO
type AlertRuleDto struct {

	// **参数解释**： 规则类型。 **约束限制**： 不涉及。 **取值范围**： - CONCURRENCY：并发数。 - FAIL_COUNT：失败次数。 - QUEUE_BACKLOG：队列积压。 **默认取值**： 不涉及。
	RuleType *string `json:"ruleType,omitempty"`

	// **参数解释**： 阈值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ThresholdValue *int32 `json:"thresholdValue,omitempty"`

	// **参数解释**： 严重级别。 **约束限制**： 不涉及。 **取值范围**： - GENERAL：一般。 - WARNING：警告。 - MAJOR：严重。 **默认取值**： 不涉及。
	Severity *string `json:"severity,omitempty"`

	// **参数解释**： 是否启用。 **约束限制**： 不涉及。 **取值范围**： - true：启用。 - false：禁用。 **默认取值**： 不涉及。
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

func (o AlertRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleDto struct{}"
	}

	return strings.Join([]string{"AlertRuleDto", string(data)}, " ")
}
