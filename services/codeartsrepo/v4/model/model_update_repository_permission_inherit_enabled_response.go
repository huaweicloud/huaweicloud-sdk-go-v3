package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryPermissionInheritEnabledResponse Response Object
type UpdateRepositoryPermissionInheritEnabledResponse struct {

	// **参数解释：** 权限继承设置。 **约束限制：** 不涉及。 **取值范围：** - true，使用上层权限配置，如果上层是代码组使用代码组权限配置，如果上层是项目使用项目权限配置。 - false，使用仓库权限配置。
	InheritParentPermission *bool `json:"inherit_parent_permission,omitempty"`
	HttpStatusCode          int   `json:"-"`
}

func (o UpdateRepositoryPermissionInheritEnabledResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryPermissionInheritEnabledResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryPermissionInheritEnabledResponse", string(data)}, " ")
}
