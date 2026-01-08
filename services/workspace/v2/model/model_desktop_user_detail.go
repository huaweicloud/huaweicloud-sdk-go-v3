package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopUserDetail 用户详细信息。
type DesktopUserDetail struct {

	// 用户id。
	Id *string `json:"id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户所属域，domain为空时，默认主域。
	Domain *string `json:"domain,omitempty"`

	// 邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 权限组。
	PermissionGroup *string `json:"permission_group,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面ip。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`
}

func (o DesktopUserDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopUserDetail struct{}"
	}

	return strings.Join([]string{"DesktopUserDetail", string(data)}, " ")
}
