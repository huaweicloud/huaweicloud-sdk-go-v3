package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CapacitySettingVo struct {

	// **参数解释**: 容量阈值。 **取值范围**: 不涉及。
	CapacityThreshold *int32 `json:"capacity_threshold,omitempty"`

	// **参数解释**: 消息类型。 **取值范围**: 不涉及。
	MessageType *int32 `json:"message_type,omitempty"`

	// **参数解释**: 是否启用邮件。 **取值范围**: 不涉及。
	IsMailEnabled *int32 `json:"is_mail_enabled,omitempty"`

	// **参数解释**: 是否启用短信。 **取值范围**: 不涉及。
	IsSmsEnabled *int32 `json:"is_sms_enabled,omitempty"`
}

func (o CapacitySettingVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacitySettingVo struct{}"
	}

	return strings.Join([]string{"CapacitySettingVo", string(data)}, " ")
}
