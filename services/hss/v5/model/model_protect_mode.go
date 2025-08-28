package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectMode **参数解释**: 防护模式 **约束限制**: 不涉及 **取值范围**: - recovery ：拦截模式。 - alarm ：告警模式。  **默认取值**: recovery
type ProtectMode struct {
}

func (o ProtectMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectMode struct{}"
	}

	return strings.Join([]string{"ProtectMode", string(data)}, " ")
}
