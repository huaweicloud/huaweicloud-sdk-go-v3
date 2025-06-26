package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsResponse Response Object
type ListSnapshotsResponse struct {

	// **参数解释**： 快照对象列表。 **取值范围**： 不涉及。
	Snapshots *[]Snapshots `json:"snapshots,omitempty"`

	// **参数解释**： 列表总数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsResponse", string(data)}, " ")
}
