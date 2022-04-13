package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNotificationsResponse struct {
	// 通知列表

	Notifications  *[]NotificationList `json:"notifications,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationsResponse", string(data)}, " ")
}
