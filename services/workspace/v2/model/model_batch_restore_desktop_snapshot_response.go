package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestoreDesktopSnapshotResponse Response Object
type BatchRestoreDesktopSnapshotResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRestoreDesktopSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreDesktopSnapshotResponse struct{}"
	}

	return strings.Join([]string{"BatchRestoreDesktopSnapshotResponse", string(data)}, " ")
}
