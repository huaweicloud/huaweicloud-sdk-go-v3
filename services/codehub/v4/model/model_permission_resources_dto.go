package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionResourcesDto struct {

	// **参数解释：** 资源id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 资源中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：** 禁用资源名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ResourceNameDisplay *string `json:"resource_name_display,omitempty"`

	// **参数解释：** 禁用资源中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ResourceNameCnDisplay *string `json:"resource_name_cn_display,omitempty"`

	// **参数解释：** 权限路径。 **取值范围：** 字符串长度不少于1，不超过1000。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	Scope *string `json:"scope,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o PermissionResourcesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionResourcesDto struct{}"
	}

	return strings.Join([]string{"PermissionResourcesDto", string(data)}, " ")
}
