package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendDesktopPoolNotificationsReq 发送桌面池消息通知请求体。
type SendDesktopPoolNotificationsReq struct {

	// 消息通知内容。
	Notifications string `json:"notifications"`
}

func (o SendDesktopPoolNotificationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDesktopPoolNotificationsReq struct{}"
	}

	return strings.Join([]string{"SendDesktopPoolNotificationsReq", string(data)}, " ")
}
