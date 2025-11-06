package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NamespaceBasicDto struct {

	// **参数解释：** 命名空间ID。 **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 命名空间名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 路径。 **约束限制：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 开发模式。 **约束限制：** 不涉及。
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 类型。 **约束限制：** 不涉及。
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** 完整路径。 **约束限制：** 不涉及。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释：** 完整名称。 **约束限制：** 不涉及。
	FullName *string `json:"full_name,omitempty"`

	// **参数解释：** 父级ID。 **约束限制：** 不涉及。
	ParentId *int32 `json:"parent_id,omitempty"`

	// **参数解释：** 可见级别。 **约束限制：** 不涉及。
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// **参数解释：** 开启文件权限控制。 **约束限制：** 不涉及。
	EnableFileControl *bool `json:"enable_file_control,omitempty"`

	// **参数解释：** 所有人ID。 **约束限制：** 不涉及。
	OwnerId *int32 `json:"owner_id,omitempty"`
}

func (o NamespaceBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceBasicDto struct{}"
	}

	return strings.Join([]string{"NamespaceBasicDto", string(data)}, " ")
}
