package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopSnapshotReq 批量删除快照请求体。
type BatchDeleteDesktopSnapshotReq struct {

	// 快照id数组，最多支持100。
	SnapshotIds *[]string `json:"snapshot_ids,omitempty"`
}

func (o BatchDeleteDesktopSnapshotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopSnapshotReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopSnapshotReq", string(data)}, " ")
}
