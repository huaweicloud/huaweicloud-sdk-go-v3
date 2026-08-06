package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeyPolicyResponseBodyPolicyValidityPeriod **参数解释：** 密钥策略有效期 **取值范围：** 不涉及
type ListKeyPolicyResponseBodyPolicyValidityPeriod struct {

	// **参数解释：** 密钥策略生效时间 **取值范围：** 不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释：** 密钥策略截止时间 **取值范围：** 不涉及
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListKeyPolicyResponseBodyPolicyValidityPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeyPolicyResponseBodyPolicyValidityPeriod struct{}"
	}

	return strings.Join([]string{"ListKeyPolicyResponseBodyPolicyValidityPeriod", string(data)}, " ")
}
