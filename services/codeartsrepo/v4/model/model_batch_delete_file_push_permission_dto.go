package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteFilePushPermissionDto struct {

	// **参数解释：** 推送权限ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Ids *[]int32 `json:"ids,omitempty"`
}

func (o BatchDeleteFilePushPermissionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteFilePushPermissionDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteFilePushPermissionDto", string(data)}, " ")
}
