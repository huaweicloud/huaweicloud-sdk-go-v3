package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotificationSettingRequest Request Object
type ShowNotificationSettingRequest struct {

	// 消息通知配置的唯一标识符。
	NotificationSettingId string `json:"notification_setting_id"`
}

func (o ShowNotificationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowNotificationSettingRequest", string(data)}, " ")
}
