package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckSnapshotReq struct {

	// **参数解释**： 快照名称。 **约束限制**： 不涉及。 **取值范围**： 非空字符串。 **默认取值**： 不涉及。
	SnapshotName *string `json:"snapshot_name,omitempty"`
}

func (o CheckSnapshotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSnapshotReq struct{}"
	}

	return strings.Join([]string{"CheckSnapshotReq", string(data)}, " ")
}
