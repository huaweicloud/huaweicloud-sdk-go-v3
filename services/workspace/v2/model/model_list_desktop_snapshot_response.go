package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopSnapshotResponse Response Object
type ListDesktopSnapshotResponse struct {

	// 桌面快照详情列表。
	DesktopSnapshots *[]DesktopSnapshotDetailInfo `json:"desktop_snapshots,omitempty"`

	// 桌面快照总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopSnapshotResponse", string(data)}, " ")
}
