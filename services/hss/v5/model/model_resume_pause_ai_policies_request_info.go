package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResumePauseAiPoliciesRequestInfo struct {

	// **参数解释**: 是否启用 **约束限制**: 必填 **取值范围**: - false：否 - true：是  **默认取值**: 不涉及
	Enabled bool `json:"enabled"`

	// **参数解释**: 策略ID **约束限制**: 必填 **取值范围**: 字符长度1-20位 **默认取值**: 不涉及
	PolicyId string `json:"policy_id"`
}

func (o ResumePauseAiPoliciesRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPoliciesRequestInfo struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPoliciesRequestInfo", string(data)}, " ")
}
