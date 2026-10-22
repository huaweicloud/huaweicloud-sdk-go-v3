package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiGuardPolicyRequest Request Object
type DeleteAiGuardPolicyRequest struct {

	// **参数解释：** 策略 ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`
}

func (o DeleteAiGuardPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiGuardPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAiGuardPolicyRequest", string(data)}, " ")
}
