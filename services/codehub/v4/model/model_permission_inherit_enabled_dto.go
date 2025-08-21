package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionInheritEnabledDto 权限继承设置请求体
type PermissionInheritEnabledDto struct {

	// **参数解释：** 权限继承设置。 **约束限制：** 不涉及。 **取值范围：** - true，使用上层权限配置，如果上层是代码组使用代码组权限配置，如果上层是项目使用项目权限配置。 - false，使用仓库权限配置。
	InheritParentPermission bool `json:"inherit_parent_permission"`
}

func (o PermissionInheritEnabledDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionInheritEnabledDto struct{}"
	}

	return strings.Join([]string{"PermissionInheritEnabledDto", string(data)}, " ")
}
