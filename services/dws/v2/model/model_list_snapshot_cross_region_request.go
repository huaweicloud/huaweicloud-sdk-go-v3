package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionRequest Request Object
type ListSnapshotCrossRegionRequest struct {

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSnapshotCrossRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionRequest", string(data)}, " ")
}
