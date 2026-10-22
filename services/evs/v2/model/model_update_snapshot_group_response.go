package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSnapshotGroupResponse Response Object
type UpdateSnapshotGroupResponse struct {
	SnapshotGroup  *SnapshotGroupDetail `json:"snapshot_group,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o UpdateSnapshotGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSnapshotGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateSnapshotGroupResponse", string(data)}, " ")
}
