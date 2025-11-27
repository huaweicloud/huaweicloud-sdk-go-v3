package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WtpProtectMode **参数解释**: 防护模式 **约束限制**: 不涉及 **取值范围**: - alarm：告警模式 - block：拦截模式  **默认取值**: alarm
type WtpProtectMode struct {
}

func (o WtpProtectMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WtpProtectMode struct{}"
	}

	return strings.Join([]string{"WtpProtectMode", string(data)}, " ")
}
