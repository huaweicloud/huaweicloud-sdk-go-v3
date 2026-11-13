package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopUsernameReq 更换桌面关联用户名请求。
type UpdateDesktopUsernameReq struct {

	// 桌面关联原用户名，只传用户名，不带域信息。
	OldUsername string `json:"old_username"`

	// 桌面关联原用户名的域。
	OldUserDomain *string `json:"old_user_domain,omitempty"`

	// 桌面关联新用户名，只传用户名，不带域信息。
	NewUsername string `json:"new_username"`

	// 桌面关联新用户名的域。
	NewUserDomain *string `json:"new_user_domain,omitempty"`

	// 桌面关联新用户名后是否重启虚拟机，默认为true，表示需要重启虚拟机。若用户此前已登录过云桌面，建议设置为true，以确保会话资源彻底释放、避免旧会话残留占用，并保障用户名修改后能正常接入新桌面。
	IsReboot *bool `json:"is_reboot,omitempty"`
}

func (o UpdateDesktopUsernameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopUsernameReq struct{}"
	}

	return strings.Join([]string{"UpdateDesktopUsernameReq", string(data)}, " ")
}
