package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInstancesDesktopInfo 分配用户请求桌面信息。
type DetachInstancesDesktopInfo struct {

	// 分配的桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 表示解关联所有用户：true，如果指定那么detach_user_infos会失效；指定解关联用户：false，通过detach_user_infos指定解关联的用户。
	IsDetachAllUsers *bool `json:"is_detach_all_users,omitempty"`

	// 解分配的用户信息列表。
	DetachUserInfos *[]DetachInstancesUserInfo `json:"detach_user_infos,omitempty"`
}

func (o DetachInstancesDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInstancesDesktopInfo struct{}"
	}

	return strings.Join([]string{"DetachInstancesDesktopInfo", string(data)}, " ")
}
