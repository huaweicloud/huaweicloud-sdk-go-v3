package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateFilePushPermissionBodyDto struct {

	// **参数解释：** 文件推送权限更新列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Permissions *[]FilePushPermissionUpdateDto `json:"permissions,omitempty"`
}

func (o BatchUpdateFilePushPermissionBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateFilePushPermissionBodyDto struct{}"
	}

	return strings.Join([]string{"BatchUpdateFilePushPermissionBodyDto", string(data)}, " ")
}
