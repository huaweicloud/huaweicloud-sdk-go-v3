package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopDetachInfo 桌面解绑信息
type DesktopDetachInfo struct {

	// 桌面id
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称
	DesktopName *string `json:"desktop_name,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 用户权限
	UserGroup *string `json:"user_group,omitempty"`

	// 解绑时间
	DetachTime *string `json:"detach_time,omitempty"`

	// 对象类型，可选值为： - USER：用户。 - GROUP：用户组。
	Type *string `json:"type,omitempty"`
}

func (o DesktopDetachInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopDetachInfo struct{}"
	}

	return strings.Join([]string{"DesktopDetachInfo", string(data)}, " ")
}
