package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSnapshotGroupResponse Response Object
type ShowSnapshotGroupResponse struct {
	SnapshotGroup  *SnapshotGroupDetail `json:"snapshot_group,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSnapshotGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSnapshotGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowSnapshotGroupResponse", string(data)}, " ")
}
