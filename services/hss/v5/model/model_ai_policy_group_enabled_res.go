package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupEnabledRes **参数解释**: 策略组是否启用 **取值范围**: - true：已启用 - false：未启用
type AiPolicyGroupEnabledRes struct {
}

func (o AiPolicyGroupEnabledRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupEnabledRes struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupEnabledRes", string(data)}, " ")
}
