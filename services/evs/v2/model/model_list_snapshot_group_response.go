package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotGroupResponse Response Object
type ListSnapshotGroupResponse struct {

	// **参数解释** 快照一致性组列表信息。 **取值范围** 不涉及。
	SnapshotGroups *[]SnapshotGroupDetail `json:"snapshot_groups,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSnapshotGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotGroupResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotGroupResponse", string(data)}, " ")
}
