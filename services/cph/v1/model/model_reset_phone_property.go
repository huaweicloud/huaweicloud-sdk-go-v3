package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetPhoneProperty 云手机属性信息。
type ResetPhoneProperty struct {

	// 云手机id。
	PhoneId string `json:"phone_id"`

	// 云手机属性列表，为Json格式字符串。
	Property *string `json:"property,omitempty"`

	// 是否恢复出厂设置，设为true 会在重置手机的基础上，清除手机所有历史属性配置记录。
	FactoryResetEnabled *bool `json:"factory_reset_enabled,omitempty"`
}

func (o ResetPhoneProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPhoneProperty struct{}"
	}

	return strings.Join([]string{"ResetPhoneProperty", string(data)}, " ")
}
