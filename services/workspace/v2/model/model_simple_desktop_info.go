package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleDesktopInfo struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 创建时间。
	Created *string `json:"created,omitempty"`

	// 桌面ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 权限组。
	UserGroup *string `json:"user_group,omitempty"`

	// 桌面的SID信息。
	Sid *string `json:"sid,omitempty"`

	// ou名称。
	OuName *string `json:"ou_name,omitempty"`
}

func (o SimpleDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleDesktopInfo struct{}"
	}

	return strings.Join([]string{"SimpleDesktopInfo", string(data)}, " ")
}
