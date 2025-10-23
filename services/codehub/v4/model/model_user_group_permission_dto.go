package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserGroupPermissionDto current user group permission
type UserGroupPermissionDto struct {

	// **参数解释：** 是否可以创建代码组。
	CanCreateGroup *bool `json:"can_create_group,omitempty"`

	// **参数解释：** 是否可以创建仓库。
	CanCraeteProject *bool `json:"can_craete_project,omitempty"`

	// **参数解释：** 是否可以设置代码组。
	CanSetGroup *bool `json:"can_set_group,omitempty"`

	// **参数解释：** 代码组id。
	GroupId *int32 `json:"group_id,omitempty"`

	// **参数解释：** 可见性, private public。
	GroupVisibility *string `json:"group_visibility,omitempty"`
}

func (o UserGroupPermissionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserGroupPermissionDto struct{}"
	}

	return strings.Join([]string{"UserGroupPermissionDto", string(data)}, " ")
}
