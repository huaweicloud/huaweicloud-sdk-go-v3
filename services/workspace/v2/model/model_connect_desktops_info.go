package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectDesktopsInfo 桌面登录状态信息。
type ConnectDesktopsInfo struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面的登录状态。  - UNREGISTER：表示桌面未注册时的状态（桌面启动后，会自动注册）。关机后也会出现未注册的状态。 - REGISTERED：表示桌面注册以后，等待用户连接的状态。 - CONNECTED：表示用户已经成功登录，正在使用桌面。 - DISCONNECTED：表示桌面与客户端断开会话后显示的状态，可能为关闭客户端窗口，或客户端与桌面网络断开引起。
	ConnectStatus *string `json:"connect_status,omitempty"`

	// 桌面已分配的用户或用户组信息列表。
	AttachUsers *[]AttachUsersInfo `json:"attach_users,omitempty"`
}

func (o ConnectDesktopsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectDesktopsInfo struct{}"
	}

	return strings.Join([]string{"ConnectDesktopsInfo", string(data)}, " ")
}
