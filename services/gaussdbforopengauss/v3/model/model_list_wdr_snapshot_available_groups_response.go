package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWdrSnapshotAvailableGroupsResponse Response Object
type ListWdrSnapshotAvailableGroupsResponse struct {

	// **参数解释**: 分组数量。 **取值范围**: 不涉及。
	TotalCount *int64 `json:"total_count,omitempty"`

	// **参数解释**: 分组记录列表。
	Groups         *[]WdrSnapshotGroupResult `json:"groups,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListWdrSnapshotAvailableGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWdrSnapshotAvailableGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListWdrSnapshotAvailableGroupsResponse", string(data)}, " ")
}
