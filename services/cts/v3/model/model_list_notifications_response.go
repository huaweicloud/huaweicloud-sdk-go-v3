package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationsResponse Response Object
type ListNotificationsResponse struct {

	// 关键操作通知列表。
	Notifications  *[]NotificationsResponseBody `json:"notifications,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationsResponse struct{}"
	}

	return strings.Join([]string{"ListNotificationsResponse", string(data)}, " ")
}
