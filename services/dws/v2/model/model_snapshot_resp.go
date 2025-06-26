package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotResp **参数解释**： 创建快照响应。 **取值范围**： 不涉及。
type SnapshotResp struct {

	// **参数解释**： 快照ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o SnapshotResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotResp struct{}"
	}

	return strings.Join([]string{"SnapshotResp", string(data)}, " ")
}
