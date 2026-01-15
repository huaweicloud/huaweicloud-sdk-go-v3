package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiOpsRequestBodyAlarm **参数解释**： 检测报告发送，当前功能已废弃。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type CreateAiOpsRequestBodyAlarm struct {

	// **参数解释**： 报告发送风险类别，当前功能已废弃。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释**： 报告发送主题，当前功能已废弃。 **约束限制**： 不涉及 **默认取值**： 不涉及
	SmnTopic *string `json:"smn_topic,omitempty"`
}

func (o CreateAiOpsRequestBodyAlarm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiOpsRequestBodyAlarm struct{}"
	}

	return strings.Join([]string{"CreateAiOpsRequestBodyAlarm", string(data)}, " ")
}
