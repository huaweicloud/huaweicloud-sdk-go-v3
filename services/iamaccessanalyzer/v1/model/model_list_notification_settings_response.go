package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationSettingsResponse Response Object
type ListNotificationSettingsResponse struct {

	// 消息通知配置列表。
	NotificationSettings *[]NotificationSettingSummary `json:"notification_settings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNotificationSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationSettingsResponse", string(data)}, " ")
}
