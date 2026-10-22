package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardPolicyRequest Request Object
type UpdateAiGuardPolicyRequest struct {

	// **参数解释：** 策略 ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	Body *UpdateAiGuardPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAiGuardPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardPolicyRequest", string(data)}, " ")
}
