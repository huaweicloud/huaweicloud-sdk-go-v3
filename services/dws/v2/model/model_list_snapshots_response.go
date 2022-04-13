package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSnapshotsResponse struct {
	// 快照对象列表

	Snapshots *[]Snapshots `json:"snapshots,omitempty"`
	// 快照对象列表总数

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
