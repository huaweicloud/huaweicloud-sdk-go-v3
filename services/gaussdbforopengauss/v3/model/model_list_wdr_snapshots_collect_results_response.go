package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWdrSnapshotsCollectResultsResponse Response Object
type ListWdrSnapshotsCollectResultsResponse struct {

	// **参数解释**： 总记录数。 **取值范围**： 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**： WDR快照信息列表。
	WdrSnapshots   *[]CollectedWdrSnapshotInfoResult `json:"wdr_snapshots,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListWdrSnapshotsCollectResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWdrSnapshotsCollectResultsResponse struct{}"
	}

	return strings.Join([]string{"ListWdrSnapshotsCollectResultsResponse", string(data)}, " ")
}
