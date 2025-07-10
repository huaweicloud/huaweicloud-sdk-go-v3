package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendDesktopPoolNotificationsResponse Response Object
type SendDesktopPoolNotificationsResponse struct {

	// 发送桌面消息任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SendDesktopPoolNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDesktopPoolNotificationsResponse struct{}"
	}

	return strings.Join([]string{"SendDesktopPoolNotificationsResponse", string(data)}, " ")
}
