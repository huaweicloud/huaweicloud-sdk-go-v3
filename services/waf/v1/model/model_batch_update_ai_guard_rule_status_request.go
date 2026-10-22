package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAiGuardRuleStatusRequest Request Object
type BatchUpdateAiGuardRuleStatusRequest struct {

	// **参数解释：** 策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	Body *BatchUpdateAiGuardRuleStatusRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateAiGuardRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAiGuardRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateAiGuardRuleStatusRequest", string(data)}, " ")
}
