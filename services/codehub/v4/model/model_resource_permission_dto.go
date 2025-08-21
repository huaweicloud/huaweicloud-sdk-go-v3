package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcePermissionDto struct {

	// **参数解释：** 记录id。
	PermissionId *int32 `json:"permission_id,omitempty"`

	// **参数解释：** 操作。 **取值范围：** 字符串长度不少于1，不超过1000。
	Action *string `json:"action,omitempty"`

	// **参数解释：** 操作名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释：** 操作中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	DisplayNameCn *string `json:"display_name_cn,omitempty"`

	// **参数解释：** 是否开启。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 是否编辑。
	Editable *bool `json:"editable,omitempty"`
}

func (o ResourcePermissionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePermissionDto struct{}"
	}

	return strings.Join([]string{"ResourcePermissionDto", string(data)}, " ")
}
