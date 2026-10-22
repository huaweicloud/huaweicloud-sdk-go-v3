package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSnapshotGroupResponseBody struct {

	// **参数解释** 快照一致性组名称。 **约束限制** 最大支持255个字节。 **取值范围** 最小长度：0 最大长度：255 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 快照一致性组描述。 **约束限制** 最大支持255个字节。 **取值范围** 最小长度：0 最大长度：255 **默认取值** 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o UpdateSnapshotGroupResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotGroupResponseBody struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotGroupResponseBody", string(data)}, " ")
}
