package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationSettingUrn 消息通知配置的唯一资源标识符。
type NotificationSettingUrn struct {
}

func (o NotificationSettingUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationSettingUrn struct{}"
	}

	return strings.Join([]string{"NotificationSettingUrn", string(data)}, " ")
}
