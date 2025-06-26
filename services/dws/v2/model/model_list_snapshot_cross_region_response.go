package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionResponse Response Object
type ListSnapshotCrossRegionResponse struct {

	// **参数解释**： 区域列表。 **取值范围**： 不涉及。
	Regions *[]SnapshotRegion `json:"regions,omitempty"`

	// **参数解释**： 总数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotCrossRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionResponse", string(data)}, " ")
}
