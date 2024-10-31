package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserPrivilegeGroupReq 批量修改用户权限组请求。
type ChangeUserPrivilegeGroupReq struct {

	// 桌面信息列表。
	Desktops []ChangeUserPrivilegeGroupDesktopInfo `json:"desktops"`

	// 桌面用户所属的用户权限组。desktops参数中，各个桌面所属的未传的桌面用户将应用该权限组。 - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserPrivilegeGroup *string `json:"user_privilege_group,omitempty"`
}

func (o ChangeUserPrivilegeGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserPrivilegeGroupReq struct{}"
	}

	return strings.Join([]string{"ChangeUserPrivilegeGroupReq", string(data)}, " ")
}
