package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopRequest Request Object
type DeleteDesktopRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`

	// 删除桌面后，如果当前用户没有其它桌面，可以删除桌面用户。true：删除用户，false：不删除用户，默认为false。
	DeleteUsers *bool `json:"delete_users,omitempty"`

	// 删除桌面后，是否给桌面用户发送系统通知邮件。true：发送，false：不发送。默认为true。
	EmailNotification *bool `json:"email_notification,omitempty"`
}

func (o DeleteDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesktopRequest", string(data)}, " ")
}
