package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestoreDesktopSnapshotRequest Request Object
type BatchRestoreDesktopSnapshotRequest struct {
	Body *BatchRestoreSnapshotReq `json:"body,omitempty"`
}

func (o BatchRestoreDesktopSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreDesktopSnapshotRequest struct{}"
	}

	return strings.Join([]string{"BatchRestoreDesktopSnapshotRequest", string(data)}, " ")
}
