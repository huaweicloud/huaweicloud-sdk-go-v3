package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationSettingRequest Request Object
type DeleteNotificationSettingRequest struct {

	// 消息通知配置的唯一标识符。
	NotificationSettingId string `json:"notification_setting_id"`
}

func (o DeleteNotificationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationSettingRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotificationSettingRequest", string(data)}, " ")
}
