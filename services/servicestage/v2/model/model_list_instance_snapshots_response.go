package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstanceSnapshotsResponse struct {

	// 快照总数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 快照列表。
	Snapshots      *[]InstanceSnapshotView `json:"snapshots,omitempty" xml:"snapshots"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListInstanceSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceSnapshotsResponse", string(data)}, " ")
}
