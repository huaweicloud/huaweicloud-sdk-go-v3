package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserPrivilegeGroupDesktopInfo 修改桌面用户权限请求桌面信息。
type ChangeUserPrivilegeGroupDesktopInfo struct {

	// 待修改的桌面ID。
	DesktopId string `json:"desktop_id"`

	// 待修改用户权限组的用户信息。未传该参数的桌面的用于将应用上层结构的参数user_privilege_group作为其用户的用户权限组。
	AttachUserInfos *[]ChangeUserPrivilegeGroupUserInfo `json:"attach_user_infos,omitempty"`
}

func (o ChangeUserPrivilegeGroupDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserPrivilegeGroupDesktopInfo struct{}"
	}

	return strings.Join([]string{"ChangeUserPrivilegeGroupDesktopInfo", string(data)}, " ")
}
