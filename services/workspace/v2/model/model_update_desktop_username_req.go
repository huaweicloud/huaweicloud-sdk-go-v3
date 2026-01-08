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

	// 桌面关联新用户名后是否重启虚拟机，默认不重启。
	IsReboot *bool `json:"is_reboot,omitempty"`
}

func (o UpdateDesktopUsernameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopUsernameReq struct{}"
	}

	return strings.Join([]string{"UpdateDesktopUsernameReq", string(data)}, " ")
}
