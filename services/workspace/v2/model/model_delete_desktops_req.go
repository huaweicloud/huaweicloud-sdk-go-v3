package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量删除桌面请求。
type DeleteDesktopsReq struct {

	// 待删除的桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`

	// 删除桌面后，如果当前用户没有其它桌面，可以删除桌面用户。true：删除用户，false：不删除用户，默认为false。
	DeleteUsers *bool `json:"delete_users,omitempty"`

	// 是否邮件通知，true：邮件通知，false：不通知，默认值true。
	EmailNotification *bool `json:"email_notification,omitempty"`
}

func (o DeleteDesktopsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopsReq struct{}"
	}

	return strings.Join([]string{"DeleteDesktopsReq", string(data)}, " ")
}
