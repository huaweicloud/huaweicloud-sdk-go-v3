package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotRegion **参数解释**： 快照区域信息。 **取值范围**： 不涉及。
type SnapshotRegion struct {

	// **参数解释**： 区域ID。 **取值范围**： 不涉及。
	RegionId *string `json:"region_id,omitempty"`
}

func (o SnapshotRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotRegion struct{}"
	}

	return strings.Join([]string{"SnapshotRegion", string(data)}, " ")
}
