package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestoreSnapshotReq 批量恢复快照请求体。
type BatchRestoreSnapshotReq struct {

	// 桌面快照记录id数组，最多支持100。
	SnapshotIds *[]string `json:"snapshot_ids,omitempty"`
}

func (o BatchRestoreSnapshotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreSnapshotReq struct{}"
	}

	return strings.Join([]string{"BatchRestoreSnapshotReq", string(data)}, " ")
}
