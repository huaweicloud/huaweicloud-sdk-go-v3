package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotsDetailsResponse Response Object
type ListSnapshotsDetailsResponse struct {

	// **参数解释** 快照信息。 **取值范围** 不涉及。
	Snapshots      *[]SnapshotDetailV5 `json:"snapshots,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSnapshotsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsDetailsResponse", string(data)}, " ")
}
