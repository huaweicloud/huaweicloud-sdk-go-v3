package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationSettingId 消息通知配置的唯一标识符。
type NotificationSettingId struct {
}

func (o NotificationSettingId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationSettingId struct{}"
	}

	return strings.Join([]string{"NotificationSettingId", string(data)}, " ")
}
