package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilePushPermissionUpdateDto struct {

	// **参数解释：** 仓库文件推送权限ID **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 目录路径或通配符 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 事件设置列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Actions *[]FilePushPermissionActionBodyDto `json:"actions,omitempty"`
}

func (o FilePushPermissionUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilePushPermissionUpdateDto struct{}"
	}

	return strings.Join([]string{"FilePushPermissionUpdateDto", string(data)}, " ")
}
