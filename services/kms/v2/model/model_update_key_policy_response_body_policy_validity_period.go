package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyPolicyResponseBodyPolicyValidityPeriod 密钥策略有效期
type UpdateKeyPolicyResponseBodyPolicyValidityPeriod struct {

	// 密钥策略生效时间
	StartTime *string `json:"start_time,omitempty"`

	// 密钥策略截止时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o UpdateKeyPolicyResponseBodyPolicyValidityPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyPolicyResponseBodyPolicyValidityPeriod struct{}"
	}

	return strings.Join([]string{"UpdateKeyPolicyResponseBodyPolicyValidityPeriod", string(data)}, " ")
}
