package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendNotificationsReq 发送消息通知请求体。
type SendNotificationsReq struct {

	// 桌面列表。
	DesktopIds []string `json:"desktop_ids"`

	// 消息通知内容。
	Notifications string `json:"notifications"`
}

func (o SendNotificationsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendNotificationsReq struct{}"
	}

	return strings.Join([]string{"SendNotificationsReq", string(data)}, " ")
}
