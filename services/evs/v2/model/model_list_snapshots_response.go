package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSnapshotsResponse struct {

	// 快照的总数量，不受limi参数的影响。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 快照信息。
	Snapshots *[]SnapshotList `json:"snapshots,omitempty" xml:"snapshots"`

	// 云硬盘快照列表查询位置标记。当查询时指定limit时会返回该字段，返回该字段表示本次查询只查出了部分云硬盘快照信息。
	SnapshotsLinks *[]Link `json:"snapshots_links,omitempty" xml:"snapshots_links"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotsResponse", string(data)}, " ")
}
