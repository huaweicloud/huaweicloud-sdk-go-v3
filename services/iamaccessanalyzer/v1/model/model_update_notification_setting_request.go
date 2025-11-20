package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationSettingRequest Request Object
type UpdateNotificationSettingRequest struct {

	// 消息通知配置的唯一标识符。
	NotificationSettingId string `json:"notification_setting_id"`

	Body *UpdateNotificationSettingReqBody `json:"body,omitempty"`
}

func (o UpdateNotificationSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationSettingRequest", string(data)}, " ")
}
